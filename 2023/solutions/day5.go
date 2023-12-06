package year2023

import (
	"bartwestenenk/aoc/utils"
	"strings"
)

func GetDay5() [2]func(input string) int {
	return [2]func(input string) int{
		challenge5a,
		challenge5b,
	}
}

type almanacPage struct {
	dest    string
	entries []mapEntry
}
type mapEntry struct {
	srcStart  int
	destStart int
	r         int
}

func (page almanacPage) getDestination(source int) int {
	for _, entry := range page.entries {
		offset := source - entry.srcStart
		if offset > 0 && offset <= entry.r {
			return entry.destStart + offset
		}
	}
	return source
}

// Takes in a source start, a range and returns a set of new ranges that are converted by this almanac page
func (page almanacPage) getDestinationRange(sourceStart int, r int) [][2]int {
	sourceEnd := sourceStart + r

	var ranges [][2]int
	for _, entry := range page.entries {
		// Situation 1
		if sourceStart < entry.srcStart && entry.srcStart < sourceEnd && sourceEnd < (entry.srcStart+entry.r) {
			ranges = append(ranges, [2]int{entry.destStart, sourceEnd - entry.srcStart})
			continue
		}

		// Situation 2
		if sourceStart < entry.srcStart && entry.srcStart < (entry.srcStart+entry.r) && (entry.srcStart+entry.r) < sourceEnd {
			ranges = append(ranges, [2]int{entry.destStart, entry.r})
			continue
		}

		// Situation 3
		if entry.srcStart < sourceStart && sourceStart < (entry.srcStart+entry.r) && (entry.srcStart+entry.r) < sourceEnd {
			offset := sourceStart - entry.srcStart
			ranges = append(ranges, [2]int{entry.destStart + offset, entry.r - offset})
			continue
		}

		// Situation 4
		if entry.srcStart < sourceStart && sourceStart < sourceEnd && sourceEnd < (entry.srcStart+entry.r) {
			offset := sourceStart - entry.srcStart
			ranges = append(ranges, [2]int{entry.destStart + offset, r})
			continue
		}
	}

	return ranges
}

func challenge5a(input string) int {
	lines := utils.SplitLines(input)
	seeds := convertToInts(strings.Split(lines[0], ": ")[1])

	almanac := map[string]almanacPage{}

	var cSource string
	var cDestination string
	var entries []mapEntry
	for _, line := range lines[2:] {
		if line == "" {
			// Line is empty,
			// save current knowledge to almanac
			almanac[cSource] = almanacPage{
				dest:    cDestination,
				entries: entries,
			}
			// Reset knowledge
			entries = []mapEntry{}
			cSource = ""
			cDestination = ""
			continue
		}

		if cSource == "" {
			line = strings.TrimRight(line, " map:")
			splits := strings.Split(line, "-")
			cSource = splits[0]
			cDestination = splits[2]
		} else {
			inputNumbers := convertToInts(line)
			entry := mapEntry{
				srcStart:  inputNumbers[1],
				destStart: inputNumbers[0],
				r:         inputNumbers[2],
			}
			entries = append(entries, entry)
		}
	}
	almanac[cSource] = almanacPage{
		dest:    cDestination,
		entries: entries,
	}

	current := "seed"
	for current != "location" {
		currentMap := almanac[current]
		for i, seed := range seeds {
			seeds[i] = currentMap.getDestination(seed)
		}
		current = currentMap.dest
	}

	lowest := 2147483647
	for _, seed := range seeds {
		if seed < lowest {
			lowest = seed
		}
	}
	return lowest
}

func challenge5b(input string) int {
	lines := utils.SplitLines(input)

	almanac := map[string]almanacPage{}

	var cSource string
	var cDestination string
	var entries []mapEntry
	for _, line := range lines[2:] {
		if line == "" {
			// Line is empty,
			// save current knowledge to almanac
			almanac[cSource] = almanacPage{
				dest:    cDestination,
				entries: entries,
			}
			// Reset knowledge
			entries = []mapEntry{}
			cSource = ""
			cDestination = ""
			continue
		}

		if cSource == "" {
			line = strings.TrimRight(line, " map:")
			splits := strings.Split(line, "-")
			cSource = splits[0]
			cDestination = splits[2]
		} else {
			inputNumbers := convertToInts(line)
			entry := mapEntry{
				srcStart:  inputNumbers[1],
				destStart: inputNumbers[0],
				r:         inputNumbers[2],
			}
			entries = append(entries, entry)
		}
	}
	almanac[cSource] = almanacPage{
		dest:    cDestination,
		entries: entries,
	}

	seedsInput := convertToInts(strings.Split(lines[0], ": ")[1])
	//var seeds []int
	var seedRanges [][2]int
	// Load seedRanges
	for i := 0; i < len(seedsInput); i = i + 2 {
		seedRanges = append(seedRanges, [2]int{seedsInput[i], seedsInput[i+1]})
	}

	current := "seed"
	for current != "location" {
		currentMap := almanac[current]
		var seedRangeNext [][2]int
		for _, seedRange := range seedRanges {
			seedRangeNext = append(seedRangeNext, currentMap.getDestinationRange(seedRange[0], seedRange[1])...)
		}
		current = currentMap.dest

		seedRanges = seedRangeNext
	}

	lowest := 2147483647
	for _, R := range seedRanges {
		if R[0] < lowest {
			lowest = R[0]
		}
	}
	return lowest
}

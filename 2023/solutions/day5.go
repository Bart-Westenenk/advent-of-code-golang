package year2023

import (
	"bartwestenenk/aoc/utils"
	"fmt"
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
	fmt.Println("Successfully loaded almanac")

	current := "seed"
	fmt.Println(seeds)
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
	fmt.Println("Successfully loaded almanac")

	seedsInput := convertToInts(strings.Split(lines[0], ": ")[1])
	var seeds []int
	var seedRange []int
	for i := 0; i < len(seedsInput); i = i + 2 {
		fmt.Println("Load next set of seeds")
		start := seedsInput[i]
		r := seedsInput[i+1]
		for j := start; j < start+r; j++ {
			seedRange = append(seedRange, j)
		}

		fmt.Println("Loaded set of seeds")

		current := "seed"
		for current != "location" {
			currentMap := almanac[current]
			for i, seed := range seedRange {
				seedRange[i] = currentMap.getDestination(seed)
			}
			current = currentMap.dest
		}

		lowest := 9223372036854775807
		for _, seed := range seedRange {
			if seed > 0 && seed < lowest {
				lowest = seed
			}
		}

		seeds = append(seeds, lowest)
		seedRange = []int{}

		fmt.Printf("Got answer for pair %v of %v: %v\n", i/2, len(seedsInput)/2, lowest)
		fmt.Printf("Current seeds array: %v\n", seeds)
	}
	fmt.Println("Loaded all seeds")

	fmt.Println(seeds)

	lowest := 2147483647
	for _, seed := range seeds {
		if seed < lowest {
			lowest = seed
		}
	}
	return lowest
}

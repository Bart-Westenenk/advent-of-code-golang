package year2023

import (
	"bartwestenenk/aoc/utils"
	"regexp"
)

func GetDay8() [2]func(input string) int {
	return [2]func(input string) int{
		challenge8a,
		challenge8b,
	}
}

//type camelMap struct {
//	start entry
//}
//
//type entry struct {
//	left  *entry
//	right *entry
//}

func parseMap(inputLines []string) map[string][2]string {
	// Map of keys to entries. Meant for lookup during parsing, can be decomposed of when done parsing
	entryStringMap := make(map[string][2]string)

	// Regex for parsing the line
	re := regexp.MustCompile("^([A-Z]{3}) = \\(([A-Z]{3}), ([A-Z]{3})\\)$")
	// Parse camelMap
	for _, line := range inputLines {
		reResult := re.FindStringSubmatch(line)
		src := reResult[1]
		leftKey := reResult[2]
		rightKey := reResult[3]
		entryStringMap[src] = [2]string{leftKey, rightKey}
		//fmt.Printf("From %s go left for %s and right for %s\n", src, leftKey, rightKey)
	}

	//return camelMap{}
	return entryStringMap
}

func challenge8a(input string) int {
	lines := utils.SplitLines(input)
	camelMap := parseMap(lines[2:])
	directions := lines[0]
	answer := 0

	// Walk the camelMap
	current := "AAA"
	for current != "ZZZ" {
		for _, dir := range directions {
			if dir == 'L' {
				current = camelMap[current][0]
			} else {
				current = camelMap[current][1]
			}
			answer++
		}
	}

	return answer
}

func challenge8b(input string) int {
	lines := utils.SplitLines(input)
	camelMap := parseMap(lines[2:])
	directions := lines[0]

	// Walk the camelMap
	var current []string

	for src, _ := range camelMap {
		if src[2] == 'A' {
			current = append(current, src)
		}
	}

	steps := make([]int, len(current))
	for i, src := range current {
		for src[2] != 'Z' {
			for _, dir := range directions {
				if dir == 'L' {
					src = camelMap[src][0]
				} else {
					src = camelMap[src][1]
				}
				steps[i]++
			}
		}
	}

	return LCM(steps[0], steps[1], steps[1:]...)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

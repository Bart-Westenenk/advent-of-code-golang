package year2022

import (
	"bartwestenenk/aoc/utils"
	"strings"
	"unicode"
)

func GetDay3() [2]func(input string) int {
	return [2]func(input string) int{
		challenge3a,
		challenge3b,
	}
}

func challenge3a(input string) int {
	rucksacks := utils.SplitLines(input)
	answer := 0
	for _, sack := range rucksacks {
		middleIndex := len(sack) / 2
		comp1 := sack[middleIndex:]
		comp2 := sack[:middleIndex]

		var duplicate int32
		for _, item := range comp1 {
			if strings.ContainsRune(comp2, item) {
				duplicate = item
				break
			}
		}
		answer += convertPriority(duplicate)
	}
	return answer
}

func convertPriority(c rune) int {
	if unicode.IsLower(c) {
		return int(c - 'a' + 1)
	} else {
		return int(c - 'A' + 27)
	}
}

func challenge3b(input string) int {
	return 0
}

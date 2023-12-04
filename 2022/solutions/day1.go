package year2022

import (
	"bartwestenenk/aoc/utils"
	"golang.org/x/exp/slices"
	"strconv"
)

func GetDay1() [2]func(input string) int {
	return [2]func(input string) int{
		challenge1a,
		challenge1b,
	}
}

func challenge1a(input string) int {
	calories := utils.SplitLines(input)
	max := 0
	sum := 0
	for _, item := range calories {
		if item == "" {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			num, _ := strconv.Atoi(item)
			sum += num
		}
	}
	return max
}

func challenge1b(input string) int {
	calories := utils.SplitLines(input)
	var elves []int
	sum := 0
	for _, item := range calories {
		if item == "" {
			elves = append(elves, sum)
			sum = 0
		} else {
			num, _ := strconv.Atoi(item)
			sum += num
		}
	}

	// Sort the list
	slices.Sort(elves)

	// Pick the last 3 items to sum
	answer := 0
	for _, elve := range elves[len(elves)-3:] {
		answer += elve
	}

	return answer
}

package year2023

import (
	"bartwestenenk/aoc/utils"
	"fmt"
	"golang.org/x/exp/slices"
)

func GetDay9() [2]func(input string) int {
	return [2]func(input string) int{
		challenge9a,
		challenge9b,
	}
}

func challenge9a(input string) int {
	answer := 0
	lines := utils.SplitLines(input)
	for _, line := range lines {
		ints := convertToInts(line)
		inc := ints[len(ints)-1] + extrapolate(ints)
		answer += inc
		fmt.Println(inc)
	}
	return answer
}

// TOO HIGH: 1992280294

func extrapolate(deltaArray []int) int {
	var newDelta []int
	allZero := slices.Max(deltaArray) == slices.Min(deltaArray)
	// If input exists of only zeroes, exterpolation will return zero
	if allZero {
		return 0
	}

	for i := len(deltaArray) - 1; i > 0; i-- {
		delta := deltaArray[i-1] - deltaArray[i]
		newDelta = append(newDelta, delta)
	}

	// Otherwise, keep exterpolating
	return newDelta[len(newDelta)-1] - extrapolate(newDelta)
}

func extrapolateBackwards(input []int) int {
	var deltaArray []int
	allZero := true
	for i := len(input) - 1; i > 0; i-- {
		delta := input[i] - input[i-1]
		deltaArray = append(deltaArray, delta)
		if delta != 0 {
			allZero = false
		}
	}

	if allZero {
		return 0
	}

	slices.Reverse(deltaArray)

	fmt.Println(input)
	// We are not to all zeroes, so keep extrapolating backwards
	return input[0] - extrapolateBackwards(deltaArray)
}

func challenge9b(input string) int {
	answer := 0
	lines := utils.SplitLines(input)
	for _, line := range lines {
		fmt.Println("NEW ENTRY")

		ints := convertToInts(line)
		//fmt.Printf("input: %v\n", ints)
		inc := ints[0] - extrapolateBackwards(ints)
		answer += inc
	}
	return answer
}

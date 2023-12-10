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

	for i := 0; i < len(deltaArray)-1; i++ {
		delta := deltaArray[i+1] - deltaArray[i]
		newDelta = append(newDelta, delta)
	}

	// Otherwise, keep exterpolating
	return newDelta[len(newDelta)-1] + extrapolate(newDelta)
}

func extrapolateBackwards(deltaArray []int) int {
	var newDelta []int
	allZero := slices.Max(deltaArray) == slices.Min(deltaArray)
	// If input exists of only zeroes, exterpolation will return zero [BASE CASE]
	if allZero {
		return 0
	}

	for i := len(deltaArray) - 1; i > 0; i-- {
		delta := deltaArray[i-1] - deltaArray[i]
		newDelta = append(newDelta, delta)
	}

	// Otherwise, keep exterpolating
	return newDelta[len(newDelta)-1] + extrapolate(newDelta)
}

func challenge9b(input string) int {
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

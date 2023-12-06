package year2023

import (
	"bartwestenenk/aoc/utils"
	"strings"
)

func GetDay6() [2]func(input string) int {
	return [2]func(input string) int{
		challenge6a,
		challenge6b,
	}
}

func challenge6a(input string) int {
	answer := 1
	lines := utils.SplitLines(input)
	times := convertToInts(strings.TrimLeft(strings.Split(lines[0], ":")[1], " "))
	distances := convertToInts(strings.TrimLeft(strings.Split(lines[1], ":")[1], " "))

	for i := range times {
		time := times[i]
		targetDist := distances[i]

		ways := 0
		for i := 0; i < time; i++ {
			speed := i
			// (Time left * speed) = our distance
			if (time-i)*speed > targetDist {
				ways++
			}
		}

		answer *= ways
	}

	return answer
}

func challenge6b(input string) int {
	input = strings.ReplaceAll(input, " ", "")
	return challenge6a(input)
}

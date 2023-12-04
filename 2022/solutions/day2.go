package year2022

import (
	"bartwestenenk/aoc/utils"
	"strings"
)

func GetDay2() [2]func(input string) int {
	return [2]func(input string) int{
		challenge2a,
		challenge2b,
	}
}

func challenge2a(input string) int {
	winMap := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	convertMap := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	scoreMap := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	answer := 0
	guide := utils.SplitLines(input)
	for _, game := range guide {
		gameDetails := strings.Split(game, " ")
		opponent := gameDetails[0]
		you := gameDetails[1]
		answer += scoreMap[you]
		if you == convertMap[opponent] {
			answer += 3
		} else if you == winMap[opponent] {
			answer += 6
		}
	}

	return answer
}

func challenge2b(input string) int {
	cMap := map[string]map[string]string{
		"X": {
			"A": "C",
			"B": "A",
			"C": "B",
		},
		"Y": {
			"A": "A",
			"B": "B",
			"C": "C",
		},
		"Z": {
			"A": "B",
			"B": "C",
			"C": "A",
		},
	}

	scoreMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	answer := 0
	guide := utils.SplitLines(input)
	for _, game := range guide {
		gameDetails := strings.Split(game, " ")
		opponent := gameDetails[0]
		you := cMap[gameDetails[1]][opponent]
		answer += scoreMap[you]

		answer += scoreMap[gameDetails[1]]
	}

	return answer
}

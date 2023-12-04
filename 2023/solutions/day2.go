package year2023

import (
	"strconv"
	"strings"
)

func GetDay2() [2]func(input string) int {
	return [2]func(input string) int{
		challenge2a,
		challenge2b,
	}
}

func challenge2a(input string) int {
	inputArray := strings.Split(input, "\n")
	knownCubesMap := map[string]int{
		"blue":  14,
		"red":   12,
		"green": 13}
	answer := 0
	for i, s := range inputArray {
		gameId := i + 1
		setsString := strings.SplitAfter(s, ": ")[1]
		var sets [][]string
		for _, set := range strings.Split(setsString, "; ") {
			sets = append(sets, strings.Split(set, ", "))
		}

		for _, set := range sets {
			for _, item := range set {
				splitItem := strings.Split(item, " ")
				amount, _ := strconv.Atoi(splitItem[0])
				if amount > knownCubesMap[splitItem[1]] {
					goto SetViolation
				}
			}
		}

		answer += gameId
	SetViolation:
	}
	return answer
}

func challenge2b(input string) int {
	inputArray := strings.Split(input, "\n")
	answer := 0

	for _, s := range inputArray {
		setsString := strings.SplitAfter(s, ": ")[1]
		var sets [][]string
		for _, set := range strings.Split(setsString, "; ") {
			sets = append(sets, strings.Split(set, ", "))
		}

		minCubesMap := map[string]int{
			"blue":  0,
			"green": 0,
			"red":   0,
		}

		for _, set := range sets {
			for _, item := range set {
				splitItem := strings.Split(item, " ")
				amount, _ := strconv.Atoi(splitItem[0])
				if minCubesMap[splitItem[1]] < amount {
					minCubesMap[splitItem[1]] = amount
				}
			}
		}
		mult := 1
		for _, i := range minCubesMap {
			mult *= i
		}

		answer += mult
	}
	return answer
}

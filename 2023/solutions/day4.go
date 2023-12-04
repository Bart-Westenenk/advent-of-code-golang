package year2023

import (
	"golang.org/x/exp/slices"
	"math"
	"strconv"
	"strings"
)

func GetDay4() [2]func(input string) int {
	return [2]func(input string) int{
		challenge4a,
		challenge4b,
	}
}

func challenge4a(input string) int {
	games := strings.Split(input, "\n")
	answer := 0
	for _, game := range games {
		gameData := strings.Split(strings.Split(game, ":")[1], "|")
		winningNumbers := convertToInts(gameData[0])
		ourNumbers := convertToInts(gameData[1])
		winningAmount := getWins(winningNumbers, ourNumbers)

		if winningAmount == 1 {
			answer += 1
		} else {
			term := math.Pow(2, float64(winningAmount-1))
			answer += int(term)
		}

	}
	return answer
}

func getWins(winningNumbers []int, ourNumbers []int) int {
	i := 0
	for _, n := range winningNumbers {
		if slices.Contains(ourNumbers, n) {
			i++
		}
	}
	return i
}

func convertToInts(input string) []int {
	var result []int
	for _, i := range strings.Split(input, " ") {
		if i != "" {
			numInt, _ := strconv.Atoi(i)
			result = append(result, numInt)
		}
	}
	return result
}

func challenge4b(input string) int {
	games := strings.Split(input, "\n")
	answer := 0

	// Note: both zero indexed, while the game on the input is one indexed
	cardToWinsMap := map[int]int{}
	copies := map[int]int{}
	for id, game := range games {
		// The copies of this game are equal to the copies we made earlier + the original
		gameCopies := copies[id] + 1
		answer += gameCopies

		gameData := strings.Split(strings.Split(game, ":")[1], "|")
		winningNumbers := convertToInts(gameData[0])
		ourNumbers := convertToInts(gameData[1])
		cardToWinsMap[id] = getWins(winningNumbers, ourNumbers)
		for i := 0; i < cardToWinsMap[id]; i++ {
			// [id] + 1 is the next card and our starting point
			copies[id+1+i] += gameCopies
		}
	}
	return answer
}

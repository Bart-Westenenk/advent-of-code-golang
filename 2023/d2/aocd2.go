package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input of challenge1
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	answer1 := challenge1(input)
	fmt.Printf("Answer part d1: %v\n", answer1)

	answer2 := challenge2(input)
	fmt.Printf("Answer part d2: %v\n", answer2)

}

func challenge1(input []string) int {
	knownCubesMap := map[string]int{
		"blue":  14,
		"red":   12,
		"green": 13}
	answer := 0
	for i, s := range input {
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

func challenge2(input []string) int {
	answer := 0

	for _, s := range input {
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

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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
	var answer []string
	for _, inputString := range input {
		var first byte
		var last byte
		for i := 0; i < len(inputString); i++ {
			// if c is a number, then its the first number to be found
			if unicode.IsDigit(rune(inputString[i])) {
				first = inputString[i]
				break
			}
		}

		for i := len(inputString) - 1; 0 <= i; i-- {
			if unicode.IsDigit(rune(inputString[i])) {
				last = inputString[i]
				break
			}
		}
		answer = append([]string(answer), string([]byte{first, last}))
	}

	// Calculate sum
	sum := 0
	for _, term := range answer {
		termInt, _ := strconv.Atoi(term)
		sum += termInt
	}

	return sum
}

func challenge2(input []string) int {
	answer := 0
	wordToNumberMap := map[string]rune{
		"one":   rune('1'),
		"two":   rune('2'),
		"three": rune('3'),
		"four":  rune('4'),
		"five":  rune('5'),
		"six":   rune('6'),
		"seven": rune('7'),
		"eight": rune('8'),
		"nine":  rune('9'),
		"d1":    rune('1'),
		"d2":    rune('2'),
		"d3":    rune('3'),
		"4":     rune('4'),
		"5":     rune('5'),
		"6":     rune('6'),
		"7":     rune('7'),
		"8":     rune('8'),
		"9":     rune('9'),
	}

	for _, s := range input {
		var first byte
		var last byte

		firstIndex := 10000000
		lastIndex := -1
		for key, val := range wordToNumberMap {
			index := strings.Index(s, key)
			if index != -1 && index < firstIndex {
				first = string(val)[0]
				firstIndex = index
			}

			index = strings.LastIndex(s, key)
			if index != -1 && index > lastIndex {
				last = string(val)[0]
				lastIndex = index
			}
		}

		out, _ := strconv.Atoi(string([]byte{first, last}))
		answer += out
	}

	return answer
}

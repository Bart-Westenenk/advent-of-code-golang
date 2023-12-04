package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	answer := 0
	width := len(input[0])
	fmt.Printf("Max width: %v\n", width)
	length := len(input)
	fmt.Printf("Max length: %v\n", length)

	var number []int32
	var coordsToCheck [][2]int

	for i := 0; i < len(input); i++ {
		s := input[i]
		for j, r := range s {
			if unicode.IsNumber(r) {
				addition := [][2]int{
					[2]int{i - 1, j + 1},
					[2]int{i, j + 1},
					[2]int{i + 1, j + 1},
					[2]int{i - 1, j},
					[2]int{i + 1, j},
					[2]int{i - 1, j - 1},
					[2]int{i, j - 1},
					[2]int{i + 1, j - 1},
				}

				number = append(number, r)
				coordsToCheck = append(coordsToCheck, addition...)
			} else {
				// If we don't have a number, skip
				if len(number) == 0 {
					continue
				}

				// We finished the number, we can now go check for the symbols
				numberInt, _ := strconv.Atoi(string(number))
				number = []int32{}

				// Sanitize the array
				var sanitizedCoords [][2]int
				for _, ints := range coordsToCheck {
					if !(ints[0] < 0 || ints[0] > length-1 || ints[1] < 0 || ints[1] > width-1) {
						sanitizedCoords = append(sanitizedCoords, ints)
					}
				}
				// Reset the coordinates
				coordsToCheck = [][2]int{}

				for _, possibleSymbol := range sanitizedCoords {
					c := rune(input[possibleSymbol[0]][possibleSymbol[1]])
					if c != '.' {
						fmt.Printf("Considering symbol: %v at %v\n", strconv.QuoteRune(c), possibleSymbol)
					}
					if c != '.' && !unicode.IsNumber(c) {
						answer += numberInt
						fmt.Printf("Add number: %v\n", numberInt)
						break
					}
				}
			}
		}
	}
	return answer
}

func challenge2(input []string) int {
	answer := 0
	width := len(input[0])
	fmt.Printf("Max width: %v\n", width)
	length := len(input)
	fmt.Printf("Max length: %v\n", length)

	var number []int32
	var coordsToCheck [][2]int

	gears := map[[2]int][]int{}

	for i := 0; i < len(input); i++ {
		s := input[i]
		for j, r := range s {
			if unicode.IsNumber(r) {
				addition := [][2]int{
					[2]int{i - 1, j + 1},
					[2]int{i, j + 1},
					[2]int{i + 1, j + 1},
					[2]int{i - 1, j},
					[2]int{i + 1, j},
					[2]int{i - 1, j - 1},
					[2]int{i, j - 1},
					[2]int{i + 1, j - 1},
				}

				number = append(number, r)
				coordsToCheck = append(coordsToCheck, addition...)
			} else {
				// If we don't have a number, skip
				if len(number) == 0 {
					continue
				}

				// We finished the number, we can now go check for the symbols
				numberInt, _ := strconv.Atoi(string(number))
				number = []int32{}

				// Sanitize the array
				var sanitizedCoords [][2]int
				for _, ints := range coordsToCheck {
					if !(ints[0] < 0 || ints[0] > length-1 || ints[1] < 0 || ints[1] > width-1) {
						sanitizedCoords = append(sanitizedCoords, ints)
					}
				}
				// Reset the coordinates
				coordsToCheck = [][2]int{}

				for _, possibleSymbol := range sanitizedCoords {
					c := rune(input[possibleSymbol[0]][possibleSymbol[1]])
					if c == '*' {
						gears[possibleSymbol] = append(gears[possibleSymbol], numberInt)
						fmt.Printf("Add number %v to gear at: %v\n", numberInt, possibleSymbol)
						break
					}
				}
			}
		}
	}

	for _, gear := range gears {
		if len(gear) == 2 {
			answer += gear[0] * gear[1]
		}
	}
	return answer
}

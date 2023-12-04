package year2023

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func GetDay3() [2]func(input string) int {
	return [2]func(input string) int{
		challenge3a,
		challenge3b,
	}
}

func challenge3a(input string) int {
	inputArray := strings.Split(input, "\n")
	answer := 0
	width := len(inputArray[0])
	fmt.Printf("Max width: %v\n", width)
	length := len(inputArray)
	fmt.Printf("Max length: %v\n", length)

	var number []int32
	var coordsToCheck [][2]int

	for i := 0; i < len(inputArray); i++ {
		s := inputArray[i]
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
					c := rune(inputArray[possibleSymbol[0]][possibleSymbol[1]])
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

func challenge3b(input string) int {
	inputArray := strings.Split(input, "\n")
	answer := 0
	width := len(inputArray[0])
	fmt.Printf("Max width: %v\n", width)
	length := len(inputArray)
	fmt.Printf("Max length: %v\n", length)

	var number []int32
	var coordsToCheck [][2]int

	gears := map[[2]int][]int{}

	for i := 0; i < len(inputArray); i++ {
		s := inputArray[i]
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
					c := rune(inputArray[possibleSymbol[0]][possibleSymbol[1]])
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

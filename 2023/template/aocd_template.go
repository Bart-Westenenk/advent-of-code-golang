package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	return answer
}

func challenge2(input []string) int {
	answer := 0
	return answer
}

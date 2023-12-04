package main

import (
	year2023 "bartwestenenk/aoc/2023"
	"bufio"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("[!] Could not load .env")
	}

	var year int
	var day int
	var part int

	var complete bool

	flag.IntVar(&year, "y", 0, "Year to run the challenge for. 0 for current year. Defaults to current year.")
	flag.IntVar(&day, "d", 0, "Day to run the challenge for. 0 for today. Defaults to today.")
	flag.IntVar(&part, "p", 0, "Part of the days challenge. 0 for both. Defaults to both.")

	flag.BoolVar(&complete, "complete", false, "Run all known solutions of a year. Defaults to false")

	var benchmark bool
	flag.BoolVar(&benchmark, "benchmark", false, "Run the challenge with benchmarking turned on. Defaults to off.")

	flag.Parse()

	// Check if complete is set together with day and parts
	if complete && (day != 0 || part != 0) {
		log.Fatal("[!] The complete flag cannot be used with the day and part flags.")
	}

	// Configure the autoloaded day, year and parts and validate input
	now := time.Now()
	if year == 0 {
		// Set year to last edition
		if now.Month() == 12 {
			year = now.Year()
		} else {
			year = now.Year() - 1
		}
	}
	if year > now.Year() {
		log.Fatal("[!] Requesting year in the future")
	}

	if !complete {
		if day == 0 {
			// If the today is not a day for the advent of code challenge, throw an error
			if now.Month() != 12 {
				log.Fatal("[!] Day is not specified and today is not a advent of code day.")
			}
			// Set day to today
			day = now.Day()
		}

		if day > 24 {
			log.Fatal("[!] Requesting day outside the challenge interval")
		}

		if year == now.Year() && now.Month() == 12 && (day > now.Day() || now.Hour() < 6) {
			log.Fatal("[!] Requesting unreleased challenge")
		}
	}

	var parts []int
	if part == 0 {
		parts = []int{0, 1}
	} else {
		parts = []int{part}
	}

	if complete {
		var totalTime int64
		totalTime = 0
		for day, solutions := range getSolutions()[year] {
			for part, _ := range solutions {
				answer, err, t := runChallenge(year, day+1, part, benchmark)
				if err != nil {
					log.Fatalf("[!] An error occurred when retrieving the answer of the day: %v\n", err)
				}
				if benchmark {
					totalTime += t
					fmt.Printf("Answer year %v day %v part %v: %v\nCalculated in %v seconds\n", year, day, part, answer, float64(t)/1000000000)
				} else {
					fmt.Printf("Answer year %v day %v part %v: %v\n", year, day, part, answer)
				}
			}

			if benchmark {
				fmt.Printf("All solutions that have been run combined took %v seconds\n", float64(totalTime)/1000000000)
			}
		}
	} else {
		var totalTime int64
		totalTime = 0
		for _, part := range parts {
			answer, err, t := runChallenge(year, day, part, benchmark)
			if err != nil {
				log.Fatalf("[!] An error occurred when retrieving the answer of the day: %v\n", err)
			}
			if benchmark {
				totalTime += t
				fmt.Printf("Answer year %v day %v part %v: %v\nCalculated in %v seconds\n", year, day, part, answer, float64(t)/1000000000)
			} else {
				fmt.Printf("Answer year %v day %v part %v: %v\n", year, day, part, answer)
			}
		}

		if benchmark {
			fmt.Printf("All solutions that have been run combined took %v seconds\n", float64(totalTime)/1000000000)
		}
	}

	//if part == 0 {
	//	answer, err := runChallenge(year, day, 0)
	//	if err != nil {
	//		log.Fatalf("[!] An error occurred when retrieving answer of the day: %v\n", err)
	//	}
	//
	//	fmt.Printf("Answer part 1: %v\n", answer)
	//
	//	answer, err = runChallenge(year, day, 1)
	//	if err != nil {
	//		log.Fatalf("[!] An error occurred when retrieving answer of the day: %v\n", err)
	//	}
	//	fmt.Printf("Answer part 2: %v\n", answer)
	//}
}

// Runs the challenge. If benchmark is true, also record the time it takes to execute the solution function.
// returns [answer] [error (if existent)] [time (-1 if benchmark == false)]
func runChallenge(year int, day int, part int, benchmark bool) (int, error, int64) {
	// Get input from the day
	input, err := getInput(year, day)
	if err != nil {
		log.Fatalf("[!] An error occurred when retrieving the AOC input: %v\n", err)
	}

	solution := getSolutions()[year][day-1][part]

	if benchmark {
		answer, deltatime := benchmarkSolution(solution, input)
		return answer, nil, deltatime
	}

	return solution(input), nil, -1
}

func benchmarkSolution(solution func(string) int, input string) (int, int64) {
	before := time.Now()
	answer := solution(input)
	after := time.Now()
	return answer, after.UnixNano() - before.UnixNano()
}

func getSolutions() map[int][][2]func(input string) int {
	return map[int][][2]func(input string) int{
		2023: year2023.GetSolutions(),
	}
}

func getInput(year int, day int) (string, error) {
	// First check if we have the input in our filesystem based cache
	fileCachePath := fmt.Sprintf("%v/input/day%v", year, day)
	file, err := os.Open(fileCachePath)
	if err == nil {
		defer file.Close()
		var input string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			input += scanner.Text() + "\n"
		}
		// Remove the last new line character
		input = strings.TrimRight(input, "\n")
		//fmt.Println("Retrieved input from filesystem cache")
		return input, nil
	} else {
		err = nil
	}

	// It is not in our cache, get it from the advent of code servers.
	//fmt.Println("Getting input from the Advent of code servers.")
	aocUrl := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)

	req, _ := http.NewRequest("GET", aocUrl, nil)
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: os.Getenv("advent_of_code_session"),
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(aocUrl)
		return "", fmt.Errorf("could not fetch the input data for today. HTTP status: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	input := string(body)
	input = strings.TrimRight(input, "\n")

	// Input is retrieved, save it to the filesystem
	file, err = os.Create(fileCachePath)
	defer file.Close()

	if err != nil {
		return "", err
	}
	_, err = file.Write([]byte(input))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

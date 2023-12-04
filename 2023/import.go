package year2023

import year2023 "bartwestenenk/aoc/2023/solutions"

func GetSolutions() [][2]func(input string) int {
	solutions := [][2]func(input string) int{
		year2023.GetDay1(),
		year2023.GetDay2(),
		year2023.GetDay3(),
		year2023.GetDay4(),
	}
	return solutions
}

package year2022

import year2022 "bartwestenenk/aoc/2022/solutions"

func GetSolutions() [][2]func(input string) int {
	solutions := [][2]func(input string) int{
		year2022.GetDay1(),
		year2022.GetDay2(),
		year2022.GetDay3(),
		year2022.GetDay4(),
	}
	return solutions
}

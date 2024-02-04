package year2023

import "bartwestenenk/aoc/utils"

func GetDay10() [2]func(input string) int {
	return [2]func(input string) int{
		challenge10a,
		challenge10b,
	}
}

func getNext(token rune, coords [2]int, previous [2]int) [2]int {
	if token == '|' {
		if coords[1] > previous[1] {
			return [2]int{coords[0], coords[1] + 1}
		} else {
			return [2]int{coords[0], coords[1] - 1}
		}
	}

	if token == '-' {
		if coords[0] > previous[0] {
			return [2]int{coords[0] + 1, coords[1]}
		} else {
			return [2]int{coords[0] - 1, coords[1]}
		}
	}

	return [2]int{0, 0}
}

func challenge10a(input string) int {
	answer := 0
	// Get all info into lines
	lines := utils.SplitLines(input)

	// Build a 2d array of the tiles
	startPosition := [2]int{0, 0}
	var grid [][]rune
	for x, line := range lines {
		var row []rune
		for y, coord := range line {
			row = append(row, coord)
			if coord == 'S' {
				startPosition = [2]int{x, y}
			}
		}
		grid = append(grid, row)
	}

	// Initialize the next array
	previous := [4][2]int{
		startPosition,
		startPosition,
		startPosition,
		startPosition,
	}
	next := [4][2]int{
		{startPosition[0] + 1, startPosition[1] + 1},
		{startPosition[0] + 1, startPosition[1] + 1},
		{startPosition[0] + 1, startPosition[1] + 1},
		{startPosition[0] + 1, startPosition[1] + 1},
	}

	// For each direction you can go to, walk the loop, until you find S.
	// If you have found S, do the amount of steps / 2
	finished := false
	steps := 0

	for !finished {
		for i, coords := range next {
			tile := grid[coords[0]][coords[1]]
			if tile == '.' {
				// Next location is ground, don't take option into consideration anymore.
				continue
			}

			nextCoord := getNext(tile, coords, previous[i])
			if (nextCoord[0] < 0 || nextCoord[1] < 0 || ) {

			}

			steps++
		}

	}

	// Get Length of the loop

	// Return the amount of steps to go halfway of that loop
	return answer
}

func challenge10b(input string) int {
	answer := 0
	return answer
}

package day4

import (
	"fmt"

	"github.com/jkorona/aoc2025/utils"
)

func Run() {
	run1()
	run2()
}

func run1() {
	lines := utils.ReadLinesFromFile("./04/input.txt")
	grid := buildGrid(lines)
	sum := 0

	for i, row := range grid {
		for j := range row {
			if checkCell(grid, i, j) {
				sum++
			}
		}
	}

	fmt.Println("Day 4, part 1:", sum)
}

func run2() {
	lines := utils.ReadLinesFromFile("./04/input.txt")
	grid := buildGrid(lines)
	partialSum := 0
	sum := 0

	for {
		partialSum = 0
		for i, row := range grid {
			for j := range row {
				if checkCell(grid, i, j) {
					partialSum++
					grid[i][j] = '.'
				}
			}
		}
		if partialSum == 0 {
			break
		}
		sum += partialSum
	}

	fmt.Println("Day 4, part 2:", sum)
}

func checkCell(grid [][]rune, i, j int) bool {
	cell := grid[i][j]

	if cell == '@' {
		ns := checkNeighbors(grid, i, j)
		if ns < 4 {
			return true
		}
	}

	return false
}

func checkNeighbors(grid [][]rune, i, j int) int {
	offsets := []struct{ di, dj int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /*{0,0},*/, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	result := 0

	for _, offset := range offsets {
		ni := i + offset.di
		nj := j + offset.dj

		inBounds := ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0])
		if inBounds && grid[ni][nj] == '@' {
			result++
		}
	}
	return result
}

func buildGrid(lines []string) [][]rune {
	matrix := make([][]rune, 0, len(lines))

	for _, line := range lines {
		matrix = append(matrix, []rune(line))
	}

	return matrix
}

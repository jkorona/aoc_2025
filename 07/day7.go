package day7

import (
	"fmt"

	"github.com/jkorona/aoc2025/utils"
)

var test = []string{
	".......S.......",
	"...............",
	".......^.......",
	"...............",
	"......^.^......",
	"...............",
	".....^.^.^.....",
	"...............",
	"....^.^...^....",
	"...............",
	"...^.^...^.^...",
	"...............",
	"..^...^.....^..",
	"...............",
	".^.^.^.^.^...^.",
	"...............",
}

func Run() {
	run1()
	run2()
}

func run1() {
	lines := utils.ReadLinesFromFile("./07/input.txt")

	head := lines[0]
	rest := lines[1:]

	beams := make(map[int]bool, len(head))

	for idx, char := range head {
		if string(char) == "S" {
			beams[idx] = true
		}
	}

	total := 0
	for lineIdx := 1; lineIdx < len(rest); lineIdx += 2 {
		line := rest[lineIdx]
		for charIdx, char := range line {
			if string(char) == "^" {
				if beams[charIdx] {
					total += 1
					beams[charIdx-1] = true
					beams[charIdx+1] = true
					beams[charIdx] = false
				}
			}
		}
	}

	fmt.Println("Day 7, part 1:", total)
}

func run2() {
	lines := utils.ReadLinesFromFile("./07/input.txt")

	head := lines[0]
	rest := lines[2:]

	var beam int

	for idx, char := range head {
		if string(char) == "S" {
			beam = idx
			break
		}
	}

	timelines := eval(rest, beam, 2)

	fmt.Println("Day 7, part 2:", timelines)
}

var lookup = make(map[int]map[int]int)

func eval(lines []string, beam int, row int) int {
	if len(lines) == 0 {
		return 1
	}

	head := lines[0]
	rest := lines[2:]

	if head[beam] == '^' {
		if lookup[row] == nil {
			lookup[row] = make(map[int]int)
		}
		if _, ok := lookup[row][beam-1]; !ok {
			lookup[row][beam-1] = eval(rest, beam-1, row+2)
		}
		if _, ok := lookup[row][beam+1]; !ok {
			lookup[row][beam+1] = eval(rest, beam+1, row+2)
		}

		res := lookup[row][beam+1] + lookup[row][beam-1]

		return res
	}

	return eval(rest, beam, row+2)
}

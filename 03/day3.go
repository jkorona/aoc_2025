package day3

import (
	"fmt"

	. "github.com/jkorona/aoc2025/utils"
)

func Run() {
	run1()
	// run2()
}

func run1() {
	fmt.Println("Day 3 - Run 1")

	lines := ReadLinesFromFile("./03/input.txt")
	fmt.Println("Read", len(lines), "lines from input.txt")
}

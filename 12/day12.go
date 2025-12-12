package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jkorona/aoc2025/utils"
)

func Run() {
	run()
}

func run() {
	input := utils.ReadFile("./12/input.txt")
	parts := strings.Split(input, "\n\n")

	gifts := make([]int, 0)
	for _, g := range parts[0 : len(parts)-1] {
		size := strings.Count(g, "#")
		gifts = append(gifts, size)
	}

	matching := 0
	for _, t := range strings.Split(parts[len(parts)-1], "\n") {
		els := strings.Split(t, ":")
		sides := strings.Split(els[0], "x")
		a, _ := strconv.Atoi(sides[0])
		b, _ := strconv.Atoi(sides[1])
		size := a * b

		exp := 0
		for idx, d := range strings.Split(strings.Trim(els[1], " "), " ") {
			num, _ := strconv.Atoi(d)
			exp += num * gifts[idx]
		}

		if exp <= size {
			matching++
		}
	}
	fmt.Println("Day 11:", matching)
}

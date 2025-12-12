package main

import (
	"fmt"
	"os"

	day1 "github.com/jkorona/aoc2025/01"
	day2 "github.com/jkorona/aoc2025/02"
	day3 "github.com/jkorona/aoc2025/03"
	day4 "github.com/jkorona/aoc2025/04"
	day5 "github.com/jkorona/aoc2025/05"
	day6 "github.com/jkorona/aoc2025/06"
	day7 "github.com/jkorona/aoc2025/07"
	day8 "github.com/jkorona/aoc2025/08"
	day9 "github.com/jkorona/aoc2025/09"

	// day10 "github.com/jkorona/aoc2025/10"
	day11 "github.com/jkorona/aoc2025/11"
	day12 "github.com/jkorona/aoc2025/12"
)

var solutions = map[string]func(){
	"1": day1.Run,
	"2": day2.Run,
	"3": day3.Run,
	"4": day4.Run,
	"5": day5.Run,
	"6": day6.Run,
	"7": day7.Run,
	"8": day8.Run,
	"9": day9.Run,
	// "10": day10.Run,
	"11": day11.Run,
	"12": day12.Run,
}

func main() {
	dayNo := os.Args[1]
	if solution, ok := solutions[dayNo]; ok {
		solution()
	} else {
		fmt.Println("Solution for day", dayNo, "not found.")
	}
}

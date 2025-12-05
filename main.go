package main

import (
	"fmt"
	"os"

	day1 "github.com/jkorona/aoc2025/01"
	day2 "github.com/jkorona/aoc2025/02"
	day3 "github.com/jkorona/aoc2025/03"
	day4 "github.com/jkorona/aoc2025/04"
	day5 "github.com/jkorona/aoc2025/05"
)

var solutions = map[string]func(){
	"1": day1.Run,
	"2": day2.Run,
	"3": day3.Run,
	"4": day4.Run,
	"5": day5.Run,
}

func main() {
	dayNo := os.Args[1]
	if solution, ok := solutions[dayNo]; ok {
		solution()
	} else {
		fmt.Println("Solution for day", dayNo, "not found.")
	}
}

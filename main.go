package main

import (
	"fmt"
	"os"

	"github.com/jkorona/aoc2025/01"
	"github.com/jkorona/aoc2025/02"
	"github.com/jkorona/aoc2025/03"
)

// create map of day packages and run main of each
var solutions = map[string]func(){
	"1": day1.Run,
	"2": day2.Run,
	"3": day3.Run,
}

func main() {
	dayNo := os.Args[1]
	if solution, ok := solutions[dayNo]; ok {
		solution()
	} else {
		fmt.Println("Solution for day", dayNo, "not found.")
	}
}
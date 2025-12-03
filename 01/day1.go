package day1

import (
	"fmt"
	"strconv"

	. "github.com/jkorona/aoc2025/utils"
)

var test = []string{
 "L68",
 "L30",
 "R48",
 "L5",
 "R60",
 "L55",
 "L1",
 "L99",
 "R14",
 "L82",
}

var test2 = []string{
	"L300",
}


func Run() {
	run1()
	run2()
}

func run1() {
	var input = ReadLinesFromFile("./01/input.txt")
	
	var currPos = 50;
	var counter = 0;

	for _, t := range input {
		var (
			dir = t[0]
			steps, _ = strconv.Atoi(t[1:])
		)
		switch dir {
		case 'L':
			currPos = (currPos - steps) % 100
			if (currPos < 0) {
				currPos = 100 + currPos
			}
		case 'R':
			currPos = (currPos + steps) % 100
		}
		
		if currPos == 0 {
			counter++
		}
	}

	fmt.Println("Number of times at starting position:", counter)
}

func run2() {
	var input = ReadLinesFromFile("./01/input.txt")
	
	var currPos = 50;
	var counter = 0;

	for _, t := range input {
		var (
			dir = t[0]
			steps, _ = strconv.Atoi(t[1:])
			newPos int
		)
		switch dir {
		case 'L':
			newPos = currPos - steps
		case 'R':
			newPos = currPos + steps
		}

		if (newPos > 0) {
			var rounds int = (newPos / 100)
			currPos = newPos % 100
			counter += rounds
		} else if (newPos < 0) {
			var rounds int = -(newPos / 100)
			if (currPos != 0) {
				rounds++
			}
			currPos = newPos % 100
			if currPos != 0 {
				currPos = 100 + currPos
			}
			counter += rounds
		} else { // 0
			currPos = 0
			counter++
		}
	}

	fmt.Println("Number of times at starting position:", counter)
}

package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
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

func readLinesFromFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string{}
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}
	}
	return lines
}

func main() {
	// run1()
	run2()
}

func run1() {
	var input = readLinesFromFile("./input_01.txt")
	
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
	var input = readLinesFromFile("./input_01.txt")
	
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

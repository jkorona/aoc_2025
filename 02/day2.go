package day2

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/jkorona/aoc2025/utils"
)

var test = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`;

func Run() {
	run1()
	run2()
}

func run1() {
	input := ReadFile("./02/input.txt")
	ranges := strings.Split(input, ",")
	sum := 0

	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])

		for i := start; i <= end; i++ {
			codeStr := fmt.Sprintf("%d", i)
			head := codeStr[:len(codeStr)/2]
			tail := codeStr[len(codeStr)/2:]

			if (head == tail) {
				sum += i
			}
		}
	}

	fmt.Println("Total sum:", sum)	
}

func hasRepeatingPatterns(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		pattern := s[0 : i+1]
		expected := strings.Repeat(pattern, len(s)/len(pattern))
		
		if (strings.EqualFold(expected, s)) {
			return true
		}
	}
	return false
}

func run2() {
	input := ReadFile("./02/input.txt")
	ranges := strings.Split(input, ",")
	sum := 0

	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])

		for i := start; i <= end; i++ {
			codeStr := fmt.Sprintf("%d", i)

			if (hasRepeatingPatterns(codeStr)) {
				sum += i
			}
		}
	}

	fmt.Println("Total sum:", sum)	
}
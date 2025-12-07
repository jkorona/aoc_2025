package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jkorona/aoc2025/utils"
)

var test = []string{
	"123 328  51 64 ",
	" 45 64  387 23 ",
	"  6 98  215 314",
	"*   +   *   +  ",
}

func Run() {
	run1()
	run2()
}

func run1() {
	lines := utils.ReadLinesFromFile("./06/input.txt")

	ops := make([]string, 0)
	last := lines[len(lines)-1]
	rest := lines[:len(lines)-1]

	for op := range strings.FieldsSeq(last) {
		ops = append(ops, string(op))
	}

	sums := make([]int, len(ops))
	for _, line := range rest {
		for index, numStr := range strings.Fields(line) {
			value, _ := strconv.Atoi(strings.TrimSpace(numStr))
			if sums[index] == 0 {
				sums[index] = value
			} else {
				switch ops[index] {
				case "+":
					sums[index] += value
				case "*":
					sums[index] *= value
				}
			}
		}
	}

	// sum all values in sums
	total := 0
	for _, v := range sums {
		total += v
	}

	fmt.Println("Day 6, part 1:", total)
}

func run2() {
	lines := utils.ReadLinesFromFile("./06/input.txt")

	ops := make([]string, 0)
	tokens := make([]int, 0)
	last := lines[len(lines)-1]
	rest := lines[:len(lines)-1]

	// parse last line into ops and tokens
	currStrike := 1
	ops = append(ops, string(last[0]))
	for index := 1; index < len(last); index++ {
		char := last[index]
		if char == ' ' {
			currStrike++
		} else {
			ops = append(ops, string(char))
			tokens = append(tokens, currStrike-1)
			currStrike = 1
		}
	}
	tokens = append(tokens, currStrike)

	m := make(map[int][]string)

	for _, line := range rest {
		caret := 0
		for colIdx, token := range tokens {
			col := line[caret : caret+token]
			caret += token + 1 // +1 for space

			if len(m[colIdx]) == 0 {
				m[colIdx] = make([]string, token)
			}
			idx := 0
			for digitIdx := len(col) - 1; digitIdx >= 0; digitIdx-- {
				m[colIdx][idx] = m[colIdx][idx] + string(col[digitIdx])
				idx++
			}
		}
	}

	sums := make([]int, len(ops))
	for opIdx, op := range ops {
		nums := m[opIdx]

		for _, numStr := range nums {
			value, _ := strconv.Atoi(strings.TrimSpace(numStr))
			if sums[opIdx] == 0 {
				sums[opIdx] = value
			} else {
				switch op {
				case "+":
					sums[opIdx] += value
				case "*":
					sums[opIdx] *= value
				}
			}
		}
	}

	// sum all values in sums
	total := 0
	for _, v := range sums {
		total += v
	}

	fmt.Println("Day 6, part 2:", total)
}

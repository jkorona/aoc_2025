package day3

import (
	"fmt"
	"strconv"

	utils "github.com/jkorona/aoc2025/utils"
)

var test = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

func Run() {
	run1()
	run2()
}

func run1() {
	lines := utils.ReadLinesFromFile("./03/input.txt")
	sum := 0

	for _, line := range lines {
		chars := []rune(line)
		var m1, m2 int
		for i := 0; i < len(chars); i++ {
			curr, _ := strconv.Atoi(string(chars[i]))

			if curr > m1 {
				if i == len(chars)-1 {
					m2 = curr
				} else {
					m1 = curr
					m2 = 0
				}
			} else if curr > m2 {
				m2 = curr
			}
		}

		concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", m1, m2))
		sum += concat
	}

	fmt.Println("Day 3, part 1:", sum)
}

func insert(value int, arr []int, index int, size int) []int {
	limit := 12
	if (size - index) < 12 {
		limit = size - index
	}

	newArr := make([]int, 0, 12)

	for i := 12 - limit; i < len(arr); i++ {
		if value > arr[i] {
			newArr = append(newArr, arr[:i]...)
			newArr = append(newArr, value)
			return newArr
		}
	}

	if len(arr) < 12 {
		return append(arr, value)
	}
	return arr
}

func run2() {
	lines := utils.ReadLinesFromFile("./03/input.txt")
	sum := 0

	for _, line := range lines {
		chars := []rune(line)
		max := make([]int, 0, 12)
		size := len(chars)

		for i := range size {
			curr, _ := strconv.Atoi(string(chars[i]))
			max = insert(curr, max, i, size)
		}

		resultStr := ""
		for _, d := range max {
			resultStr += strconv.Itoa(d)
		}

		resultNum, _ := strconv.Atoi(resultStr)
		sum += resultNum
	}

	fmt.Println("Day 3, part 2:", sum)
}

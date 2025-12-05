package day5

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/jkorona/aoc2025/utils"
)

var test = []string{
	"3-5",
	"10-14",
	"16-20",
	"12-18",
	"",
	"1",
	"5",
	"8",
	"11",
	"17",
	"32",
}

type Range struct {
	start int64
	end   int64
}

func Run() {
	run1()
	run2()
}

func run1() {
	input := utils.ReadLinesFromFile("./05/input.txt")
	index := 0
	dict := make([]Range, 0, len(input))

	for {
		if input[index] == "" {
			index++
			break
		}

		var start, end int64
		fmt.Sscanf(input[index], "%d-%d", &start, &end)
		dict = append(dict, Range{start: start, end: end})

		index++
	}

	fresh := 0
	for ; index < len(input); index++ {
		var value int64
		fmt.Sscanf(input[index], "%d", &value)

		for _, r := range dict {
			if value >= r.start && value <= r.end {
				fresh++
				break
			}
		}

	}

	fmt.Println("Day 5, part 1:", fresh)
}

func run2() {
	input := utils.ReadLinesFromFile("./05/input.txt")
	var sum int64 = 0
	var ranges []Range

	for index := 0; input[index] != ""; index++ {
		var start, end int64
		fmt.Sscanf(input[index], "%d-%d", &start, &end)

		ranges = append(ranges, Range{start: start, end: end})
	}

	ranges = mergeRanges(ranges)

	for _, r := range ranges {
		sum += r.end - r.start + 1
	}

	fmt.Println("Day 5, part 2:", sum)
}

func mergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}

	slices.SortFunc(ranges, func(r1, r2 Range) int {
		return cmp.Compare(r1.start, r2.start)
	})

	merged := make([]Range, 0, len(ranges))
	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		if ranges[i].start <= current.end+1 {
			if ranges[i].end > current.end {
				current.end = ranges[i].end
			}
		} else {
			merged = append(merged, current)
			current = ranges[i]
		}
	}

	merged = append(merged, current)

	return merged
}

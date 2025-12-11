package day11

import (
	"fmt"
	"strings"

	"github.com/jkorona/aoc2025/utils"
)

type Node struct {
	name     string
	siblings []string
}

func Run() {
	run1()
	run2()
}

func run1() {
	lines := utils.ReadLinesFromFile("./11/input.txt")
	nodes := parse(lines)

	fmt.Println("Day 11, part 1:", findNumOfPaths(nodes, "you", "you", "out"))
}

func parse(lines []string) map[string]Node {
	result := make(map[string]Node)
	for _, line := range lines {
		kv := strings.Split(line, ":")
		siblings := strings.Fields(kv[1])
		result[kv[0]] = Node{kv[0], siblings}
	}
	return result
}

func findNumOfPaths(nodes map[string]Node, from, current, to string) int {
	result := 0

	startNode := nodes[current]
	for _, sibling := range startNode.siblings {
		if sibling != from {
			if sibling == to {
				result += 1
			} else {
				result += findNumOfPaths(nodes, from, sibling, to)
			}
		}
	}

	return result
}

func run2() {
	lines := utils.ReadLinesFromFile("./11/input.txt")
	nodes := parse(lines)

	fmt.Println("Day 11, part 1:", findNumOfPathsIncluding(nodes, "svr", "svr", "out", []string{"dac", "fft"}))
}

var cache = make(map[string]int)

func findNumOfPathsIncluding(nodes map[string]Node, from, current, to string, wanted []string) int {
	key := fmt.Sprintf("%s : %v", current, wanted)
	cached, ok := cache[key]
	if ok {
		return cached
	}

	result := 0

	startNode := nodes[current]
	for _, sibling := range startNode.siblings {
		if sibling != from {
			if sibling == to {
				if len(wanted) == 0 {
					result += 1
				}
			} else {
				result += findNumOfPathsIncluding(nodes, from, sibling, to, remove(wanted, sibling))
			}
		}
	}
	cache[key] = result
	return result
}

func remove(given []string, toRemove string) []string {
	new := []string{}
	for _, item := range given {
		if item != toRemove {
			new = append(new, item)
		}
	}
	return new
}

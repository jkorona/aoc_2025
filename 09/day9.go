package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jkorona/aoc2025/utils"
)

var test = []string{
	"7,1",
	"11,1",
	"11,7",
	"9,7",
	"9,5",
	"2,5",
	"2,3",
	"7,3",
}

type Point = struct{ x, y float64 }

func Run() {
	run1()
	run2()
}

func run1() {
	lines := utils.ReadLinesFromFile("./09/input.txt") // test
	points := stringsToPoints(lines)

	max := float64(0)
	for _, p1 := range points {
		for _, p2 := range points {
			size := calcRectSize(p1, p2)
			if size > max {
				max = size
			}
		}
	}

	fmt.Println("Day 9, part 1:", int(max))
}

type Edge = struct {
	p1, p2 Point
}

func run2() {
	lines := utils.ReadLinesFromFile("./09/input.txt")
	points := stringsToPoints(lines)

	edges := []Edge{}
	for i := 0; i < len(points)-1; i++ {
		p1 := points[i]
		p2 := points[i+1]
		edges = append(edges, Edge{p1, p2})
	}

	edges = append(edges, Edge{points[0], points[len(points)-1]})
	result := float64(0)
	for i := 0; i < len(points)-1; i++ {
		for j := i; j < len(points); j++ {
			fromTile := points[i]
			toTile := points[j]
			minX, maxX := minmax(fromTile.x, toTile.x)
			minY, maxY := minmax(fromTile.y, toTile.y)
			if !intersects(minX, minY, maxX, maxY, edges) {
				area := calcRectSize(fromTile, toTile)
				if area > result {
					result = area
				}
			}
		}
	}
	fmt.Println("Day 9, part 2:", int(result))
}

func stringsToPoints(input []string) []Point {
	points := []Point{}
	for _, str := range input {
		digits := strings.Split(str, ",")
		x, _ := strconv.Atoi(digits[0])
		y, _ := strconv.Atoi(digits[1])
		points = append(points, Point{float64(x), float64(y)})
	}
	return points
}

func calcRectSize(p1, p2 Point) float64 {
	// P=a⋅b=∣x2​−x1​∣⋅∣y2​−y1​∣
	return (math.Abs(p2.x-p1.x) + 1) * (math.Abs(p2.y-p1.y) + 1)
}

func minmax(a, b float64) (float64, float64) {
	if a > b {
		return b, a
	}
	return a, b
}

func intersects(minX, minY, maxX, maxY float64, edges []Edge) bool {
	for _, edge := range edges {
		edgeMinX, edgeMaxX := minmax(edge.p1.x, edge.p2.x)
		edgeMinY, edgeMaxY := minmax(edge.p1.y, edge.p2.y)
		if minX < edgeMaxX && maxX > edgeMinX && minY < edgeMaxY && maxY > edgeMinY {
			return true
		}
	}
	return false
}

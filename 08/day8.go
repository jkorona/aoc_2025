package day8

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/jkorona/aoc2025/utils"
)

var test = []string{
	"162,817,812",
	"57,618,57",
	"906,360,560",
	"592,479,940",
	"352,342,300",
	"466,668,158",
	"542,29,236",
	"431,825,988",
	"739,650,466",
	"52,470,668",
	"216,146,977",
	"819,987,18",
	"117,168,530",
	"805,96,715",
	"346,949,466",
	"970,615,88",
	"941,993,340",
	"862,61,35",
	"984,92,344",
	"425,690,689",
}

type JunctionBox struct{ x, y, z int }
type Node struct {
	coords JunctionBox
	links  []*Node
}
type Distance = struct {
	dist float64
	from *Node
	to   *Node
}

func Run() {
	run1()
	run2()
}

func run1() {
	lines := utils.ReadLinesFromFile("./08/input.txt")
	distances := calcDistances(lines)

	connections := make([]*Node, 0)
	for i := 0; i < 2000; i += 2 {
		d := distances[i]
		connections = append(connections, d.from)

		d.from.links = append(d.from.links, d.to)
		d.to.links = append(d.to.links, d.from)
	}

	circuits := make([]int, 0)
	visited := make(map[*Node]bool)
	for _, conn := range connections {
		size := countCircuit(conn, visited)
		if size != 0 {
			circuits = append(circuits, size)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(circuits)))

	mul := circuits[0] * circuits[1] * circuits[2]

	fmt.Println("Day 8, part 1:", mul)
}

func run2() {
	lines := utils.ReadLinesFromFile("./08/input.txt")
	distances := calcDistances(lines)

	for i := 0; i < len(distances); i += 2 {
		d := distances[i]

		d.from.links = append(d.from.links, d.to)
		d.to.links = append(d.to.links, d.from)

		visited := make(map[*Node]bool)
		if countCircuit(d.from, visited) == len(lines) {
			fmt.Println("Day 8, part 2:", d.from.coords.x*d.to.coords.x)
			return
		}
	}

}

func parseToJunctionBox(input string) JunctionBox {
	values := strings.Split(input, ",")
	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])
	z, _ := strconv.Atoi(values[2])

	return JunctionBox{
		x: x, y: y, z: z,
	}
}

func calcDistances(lines []string) []Distance {
	nodes := make(map[int]*Node)

	for idx, line := range lines {
		box := parseToJunctionBox(line)
		nodes[idx] = &Node{
			coords: box,
			links:  make([]*Node, 0),
		}
	}

	distances := make([]Distance, 0)
	for nodeIdx, node := range nodes {
		for nbIdx, nb := range nodes {
			if nodeIdx == nbIdx {
				continue
			}

			dist := calcDistance(node.coords, nb.coords)

			distances = append(distances, Distance{
				dist: dist,
				from: node,
				to:   nb,
			})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist < distances[j].dist
	})

	return distances
}

func calcDistance(c1 JunctionBox, c2 JunctionBox) float64 {
	return math.Sqrt(
		math.Pow(float64(c1.x)-float64(c2.x), 2) +
			math.Pow(float64(c1.y)-float64(c2.y), 2) +
			math.Pow(float64(c1.z)-float64(c2.z), 2))
}

func countCircuit(node *Node, visited map[*Node]bool) int {
	if visited[node] {
		return 0
	}

	visited[node] = true

	result := 1
	for _, next := range node.links {
		result += countCircuit(next, visited)
	}

	return result
}

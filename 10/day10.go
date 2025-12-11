package day10

import (
	"fmt"
	"strings"

	"github.com/draffensperger/golp"
	"github.com/jkorona/aoc2025/utils"
)

var test = []string{
	"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
	"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
	"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
}

type Diagram = struct {
	lights   []bool
	buttons  [][]int
	joltages []int
}

type State[T any] struct {
	clicks int
	state  T
}

func Run() {
	run1()
	run2()
}

func run1() {
	lines := utils.ReadLinesFromFile("./10/input.txt")
	diagrams := parse(lines)

	results := []int{}
	for _, diagram := range diagrams {
		results = append(results, resolveLights(diagram))
	}

	fmt.Println("Day 10, part 1:", utils.Sum(results))
}

func run2() {
	lines := utils.ReadLinesFromFile("./10/input.txt")
	diagrams := parse(lines)

	results := []int{}
	for _, diagram := range diagrams {
		results = append(results, resolveJoltages(diagram))
	}

	fmt.Println("Day 10, part 2:", utils.Sum(results))
}

func parse(lines []string) []Diagram {
	diagrams := []Diagram{}
	for _, line := range lines {
		parts := strings.Split(line, " ")

		lights := make([]bool, 0)
		for i := 1; i < len(parts[0])-1; i++ {
			value := parts[0][i] == '#'
			lights = append(lights, value)
		}

		buttons := make([][]int, len(parts)-2)
		for i := 1; i < len(parts)-1; i++ {
			buttonNums := utils.ParseStringToIntegers(parts[i][1 : len(parts[i])-1])
			buttons[i-1] = buttonNums
		}

		js := parts[len(parts)-1]
		joltages := utils.ParseStringToIntegers(js[1 : len(js)-1])

		diagrams = append(diagrams, Diagram{lights, buttons, joltages})
	}

	return diagrams
}

func resolveLights(diagram Diagram) int {
	initial := make([]bool, len(diagram.lights))

	queue := utils.NewQueue[State[[]bool]]()
	queue.Enqueue(State[[]bool]{0, initial})

	for queue.Len() > 0 {
		currentState, _ := queue.Dequeue()
		for _, button := range diagram.buttons {
			newState := pushButtons(currentState.state, button)
			if statesEql(newState, diagram.lights) {
				return currentState.clicks + 1
			}
			queue.Enqueue(State[[]bool]{currentState.clicks + 1, newState})
		}
	}

	return 0
}

func resolveJoltages(diagram Diagram) int {
	buttonsSize := len(diagram.buttons)
	joltagesSize := len(diagram.joltages)

	lp := golp.NewLP(0, buttonsSize)
	lp.SetVerboseLevel(golp.NEUTRAL)

	objectiveCoeffs := make([]float64, buttonsSize)
	for i := range buttonsSize {
		objectiveCoeffs[i] = 1.0
	}
	lp.SetObjFn(objectiveCoeffs)

	for i := range buttonsSize {
		lp.SetInt(i, true)
		lp.SetBounds(i, 0.0, float64(2000))
	}

	for i := 0; i < joltagesSize; i++ {
		var entries []golp.Entry
		for j, btn := range diagram.buttons {
			if contains(btn, i) {
				entries = append(entries, golp.Entry{Col: j, Val: 1.0})
			}
		}
		targetValue := float64(diagram.joltages[i])
		if err := lp.AddConstraintSparse(entries, golp.EQ, targetValue); err != nil {
			panic(err)
		}
	}

	status := lp.Solve()

	if status != golp.OPTIMAL {
		return 0
	}

	solution := lp.Variables()
	totalPresses := 0
	for _, val := range solution {
		totalPresses += int(val + 0.5)
	}

	return totalPresses
}

func pushButtons(state []bool, button []int) []bool {
	newState := make([]bool, len(state))
	copy(newState, state)

	for _, toggle := range button {
		newState[toggle] = !state[toggle]
	}
	return newState
}

func statesEql(given []bool, expected []bool) bool {
	for idx := range given {
		if given[idx] != expected[idx] {
			return false
		}
	}
	return true
}

func contains(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

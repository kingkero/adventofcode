package day10

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

type Direction string

const (
	EAST  Direction = "e"
	SOUTH           = "s"
	WEST            = "w"
	NORTH           = "n"
)

type Matrix struct {
	data    [][]string
	start   []int
	visited [][]int
}

func (matrix *Matrix) findStart() {
	for i, line := range matrix.data {
		for j, field := range line {
			if field == "S" {
				matrix.start = []int{i, j}
				return
			}
		}
	}
	log.Fatal("Could not find start!")
}

func NewMatrix(lines []string) *Matrix {
	data := make([][]string, len(lines))
	for i, line := range lines {
		data[i] = strings.Split(line, "")
	}

	result := &Matrix{data, nil, nil}
	result.findStart()
	result.visited = append(result.visited, result.start)

	return result
}

func (matrix Matrix) getConnection(from Direction, i, j int) []int {
	if i < 0 || j < 0 || i >= len(matrix.data) || j >= len(matrix.data[0]) {
		return nil
	}

	switch matrix.data[i][j] {
	case "|":
		if from == NORTH {
			return []int{i + 1, j}
		}
		if from == SOUTH {
			return []int{i - 1, j}
		}
	case "-":
		if from == WEST {
			return []int{i, j + 1}
		}
		if from == EAST {
			return []int{i, j - 1}
		}
	case "L":
		if from == NORTH {
			return []int{i, j + 1}
		}
		if from == EAST {
			return []int{i - 1, j}
		}
	case "J":
		if from == WEST {
			return []int{i - 1, j}
		} else if from == NORTH {
			return []int{i, j - 1}
		}
	case "7":
		if from == SOUTH {
			return []int{i, j - 1}
		}
		if from == WEST {
			return []int{i + 1, j}
		}
	case "F":
		if from == SOUTH {
			return []int{i, j + 1}
		}
		if from == EAST {
			return []int{i + 1, j}
		}
	}

	return nil
}

type CheckSide struct {
	fromDirection Direction
	coordinates   []int
}

func (matrix Matrix) getAllConnections(i, j int) [][]int {
	checks := []CheckSide{
		{EAST, []int{i, j - 1}},
		{SOUTH, []int{i - 1, j}},
		{WEST, []int{i, j + 1}},
		{NORTH, []int{i + 1, j}},
	}

	var connections [][]int
	for _, check := range checks {
		if connection := matrix.getConnection(check.fromDirection, check.coordinates[0], check.coordinates[1]); connection != nil {
			connections = append(connections, check.coordinates)
		}
	}

	return connections
}

func getPreviousDirection(current, prev []int) Direction {
	if current[0]-prev[0] == 0 {
		if current[1]-prev[1] == 1 {
			return WEST
		}
		if current[1]-prev[1] == -1 {
			return EAST
		}
	}

	if current[0]-prev[0] == -1 {
		return SOUTH
	}

	return NORTH
}

func part01(matrix *Matrix) int {
	length := 2

	startConnections := matrix.getAllConnections(matrix.start[0], matrix.start[1])

	fmt.Printf("Start is at %v, check length from %v\n", matrix.start, startConnections[0])

	ignore := matrix.start
	previous := startConnections[0]

	matrix.visited = append(matrix.visited, previous)
	from := getPreviousDirection(previous, ignore)
	next := matrix.getConnection(from, previous[0], previous[1])

	for next != nil {
		matrix.visited = append(matrix.visited, next)
		ignore = previous
		previous = next
		from = getPreviousDirection(previous, ignore)
		next = matrix.getConnection(from, previous[0], previous[1])

		length++
	}

	// create file
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	for i := 0; i < len(matrix.data); i++ {
		val := strings.Repeat("_", len(matrix.data[i]))
		for j := 0; j < len(matrix.data[i]); j++ {
			if slices.ContainsFunc(matrix.visited, func(pos []int) bool {
				return pos[0] == i && pos[1] == j
			}) {
				left := ""
				if j > 0 {
					left = val[:j]
				}
				val = left + "X" + val[j+1:]
			}
		}
		_, err := f.WriteString(val + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	return int(math.Floor(float64(length) / 2.0))
}

func part02(matrix *Matrix) int {
	result := 0

	return result
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	matrix := NewMatrix(lines)

	return part01(matrix), part02(matrix)
}

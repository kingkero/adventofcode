package day10

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/gookit/goutil/dump"
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
	data  [][]string
	start []int
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

	result := &Matrix{data, nil}
	result.findStart()

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

	/*
		if connection := matrix.getConnection(EAST, i, j-1); connection != nil {
			// connections = append(connections, connection)
			connections = append(connections, []int{i, j - 1})
		}
		if connection := matrix.getConnection(SOUTH, i-1, j); connection != nil {
			// connections = append(connections, connection)
			connections = append(connections, []int{i - 1, j})
		}
		if connection := matrix.getConnection(WEST, i, j+1); connection != nil {
			// connections = append(connections, connection)
			connections = append(connections, []int{i, j + 1})
		}
		if connection := matrix.getConnection(NORTH, i+1, j); connection != nil {
			// connections = append(connections, connection)
			connections = append(connections, []int{i + 1, j})
		}
	*/

	return connections
}

/*
func (matrix *Matrix) getNextConnection(lookFrom, before []int) []int {
	allConnections := util.Filter(matrix.getAllConnections(lookFrom[0], lookFrom[1]), func(connection []int) bool {
		return connection[0] != before[0] || connection[1] != before[1]
	})

	if len(allConnections) > 0 {
		return allConnections[0]
	}

	return nil
}
*/

func part01(lines []string) int {
	matrix := NewMatrix(lines)
	length := 2

	startConnections := matrix.getAllConnections(matrix.start[0], matrix.start[1])

	fmt.Printf("start at %v has connections %v\n", matrix.start, startConnections)

	visited := [][]int{matrix.start}
	/*

			ignore := matrix.start
			previous := startConnections[0]
			visited = append(visited, ignore)
			visited = append(visited, previous)

			nextRight := matrix.getNextConnection(previous, ignore)

			for nextRight != nil && matrix.data[nextRight[0]][nextRight[1]] != "S" {
				fmt.Printf("%v => %v \"%v\" (length %d)\n", previous, nextRight, matrix.data[nextRight[0]][nextRight[1]], length)
				if slices.ContainsFunc(visited, func(val []int) bool {
					return val[0] == nextRight[0] && val[1] == nextRight[1]
				}) {
					fmt.Println("  already visited!")
					nextRight = nil
					continue
				}
				visited = append(visited, nextRight)
				length += 2

				ignore = previous
				previous = nextRight

				nextRight = matrix.getNextConnection(previous, ignore)

			}

		if nextRight != nil {
			length += 2
		}
	*/

	dump.P(visited)

	// create file
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	for i, line := range lines {
		val := strings.Repeat("_", len(line))
		for j := 0; j < len(line); j++ {
			if slices.ContainsFunc(visited, func(pos []int) bool {
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

func part02(lines []string) int {
	result := 0

	return result
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	return part01(lines), part02(lines)
}

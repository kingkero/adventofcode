package day10

import (
	"log"
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
	return
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

func (matrix Matrix) getAllConnections(i, j int) [][]int {
	var connections [][]int
	if connection := matrix.getConnection(EAST, i, j-1); connection != nil {
		connections = append(connections, connection)
	}
	if connection := matrix.getConnection(SOUTH, i-1, j); connection != nil {
		connections = append(connections, connection)
	}
	if connection := matrix.getConnection(WEST, i, j+1); connection != nil {
		connections = append(connections, connection)
	}
	if connection := matrix.getConnection(NORTH, i+1, j); connection != nil {
		connections = append(connections, connection)
	}

	return connections
}

func (matrix *Matrix) getNextConnection(lookFrom, before []int) []int {
	allConnections := util.Filter(matrix.getAllConnections(lookFrom[0], lookFrom[1]), func(connection []int) bool {
		return connection[0] != before[0] || connection[1] != before[1]
	})

	if len(allConnections) == 1 {
		return allConnections[0]
	}

	return nil
}

func part01(lines []string) int {
	result := 0

	matrix := NewMatrix(lines)
	dump.P("starting at", matrix.start)
	// lengths := make([]int, 2)

	startConnections := matrix.getAllConnections(matrix.start[0], matrix.start[1])
	dump.P("startConnections are", startConnections)

	nextRight := matrix.getNextConnection(startConnections[0], matrix.start)
	dump.P("nextRight is", nextRight)
	nextRight = matrix.getNextConnection(nextRight, startConnections[0])
	dump.P("nextRight is", nextRight)

	/*
		checkRight := matrix.getRightLength(startConnections[0])

		dump.P(checkRight)
	*/
	return result
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

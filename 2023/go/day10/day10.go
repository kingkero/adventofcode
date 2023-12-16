package day10

import (
	"log"
	"strings"

	"github.com/gookit/goutil/dump"
	"github.com/kingkero/adventofcode/2023/go/util"
)

type Matrix struct {
	data [][]string
}

func (matrix Matrix) findStart() (int, int) {
	for i, line := range matrix.data {
		for j, field := range line {
			if field == "S" {
				return i, j
			}
		}
	}
	log.Fatal("Could not find start!")
	return -1, -1
}

type Direction string

const (
	EAST  Direction = "e"
	SOUTH           = "s"
	WEST            = "w"
	NORTH           = "n"
)

func NewMatrix(lines []string) *Matrix {
	data := make([][]string, len(lines))
	for i, line := range lines {
		data[i] = strings.Split(line, "")
	}

	return &Matrix{data}
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
			return []int{j, i - 1}
		}
	case "L":
		if from == NORTH {
			return []int{i, j + 1}
		}
		if from == EAST {
			return []int{j - 1, i}
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
			return []int{j + 1, i}
		}
	case "F":
		if from == SOUTH {
			return []int{i, j + 1}
		}
		if from == EAST {
			return []int{j + 1, i}
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

func part01(lines []string) int {
	result := 0

	matrix := NewMatrix(lines)
	// lengths := make([]int, 2)

	i, j := matrix.findStart()

	startConnections := matrix.getAllConnections(i, j)

	firstRight := matrix.getAllConnections(startConnections[0][0], startConnections[0][1])
	dump.P("firstRight", util.Filter(firstRight, func(connection []int) bool {
		return connection[0] != i || connection[1] != j
	}))

	firstLeft := matrix.getAllConnections(startConnections[1][0], startConnections[1][1])
	dump.P("firstLeft", util.Filter(firstLeft, func(connection []int) bool {
		return connection[0] != i || connection[1] != j
	}))

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

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

func (matrix *Matrix) getConnection(from Direction, i, j int) ([]int, bool) {
	result := make([]int, 2)
	if i < 0 || j < 0 || i >= len(matrix.data) || j >= len(matrix.data[0]) {
		return result, false
	}

	switch matrix.data[i][j] {
	case "J":
		if from == WEST {
			return []int{i - 1, j}, true
		} else if from == NORTH {
			return []int{i, j - 1}, true
		}
	}

	return result, false
}

func part01(lines []string) int {
	result := 0

	matrix := NewMatrix(lines)

	i, j := matrix.findStart()
	if connection, ok := matrix.getConnection(EAST, i, j+1); ok {
		dump.P(connection)
	}
	if connection, ok := matrix.getConnection(SOUTH, i-1, j); ok {
		dump.P(connection)
	}
	if connection, ok := matrix.getConnection(WEST, i, j+1); ok {
		dump.P(connection)
	}
	if connection, ok := matrix.getConnection(NORTH, i+1, j); ok {
		dump.P(connection)
	}

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
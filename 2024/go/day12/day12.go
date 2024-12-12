package day12

import (
	"fmt"
	"strconv"

	"github.com/kingkero/adventofcode/2024/go/util"
)

type Matrix struct {
	data    [][]int32
	visited [][]bool
	rows    int
	cols    int
}

func (m *Matrix) At(row, col int) int32 {
	if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
		return 0
	}

	return m.data[row][col]
}

func (m *Matrix) String() string {
	result := ""

	for rowIndex, row := range m.data {
		for _, col := range row {
			if col < 10 {
				result += " "
			}

			if col == 0 {
				result += " "
			} else {
				result += fmt.Sprintf("%d", col)
			}

			result += " "
		}
		result += "    "
		for _, visited := range m.visited[rowIndex] {
			if visited {
				result += "X"
			} else {
				result += "_"
			}
			result += " "
		}
		result += "\n"
	}

	return result
}

var modifiers = [][]int{
	{-1, 0}, // up
	{0, -1}, // left
	{1, 0},  // down
	{0, 1},  // right
}

func Part01(input []string) string {
	input = util.Filter(input, func(s string) bool {
		return s != ""
	})

	m := &Matrix{
		rows: len(input),
		cols: len(input[0]),
	}
	m.data = make([][]int32, m.rows)
	m.visited = make([][]bool, m.rows)

	// fill matrix
	for row, line := range input {
		m.data[row] = make([]int32, m.cols)
		m.visited[row] = make([]bool, m.cols)

		for col, char := range line {
			m.data[row][col] = char
		}
	}

	result := 0

	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			visited, fences := m.getAreaAndFences(row, col, m.At(row, col))

			if visited > 0 && fences > 0 {
				result += visited * fences
			}
		}
	}

	return strconv.Itoa(result)
}

func (m *Matrix) getAreaAndFences(row, col int, plant int32) (int, int) {
	if m.visited[row][col] {
		return 0, 0
	}

	m.visited[row][col] = true
	fences := 0
	visited := 1

	for _, delta := range modifiers {
		if m.At(row+delta[0], col+delta[1]) == plant {
			subVisited, subFences := m.getAreaAndFences(row+delta[0], col+delta[1], plant)

			visited += subVisited
			fences += subFences
		} else {
			fences++
		}
	}

	return visited, fences
}

func Part02(_ []string) string {
	return ""
}

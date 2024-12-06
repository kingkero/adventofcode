package day06

import (
	"fmt"
	"strconv"

	"github.com/kingkero/adventofcode/2024/go/util"
)

type Point struct {
	Row, Col  int
	Direction uint8
}

type Matrix struct {
	Rows, Cols int
	Values     []string
	Visited    []bool
}

const (
	DirectionNorth uint8 = iota
	DirectionEast
	DirectionSouth
	DirectionWest

	AllDirections
)

const Blocker = 35 // string "#"
const Start = 94   // string "^"

var matrix *Matrix
var start *Point

func Part01(input []string) string {
	matrix = &Matrix{
		Rows: len(input),
		Cols: len(input[0]),
		Values: util.Filter(input, func(s string) bool {
			return s != ""
		}),
		Visited: make([]bool, len(input)*len(input[0])*int(AllDirections)),
	}

	current := &Point{Direction: DirectionNorth}
outer:
	for row, line := range input {
		for col, char := range line {
			if char == Start {
				current.Row = row
				current.Col = col

				break outer
			}
		}
	}
	start = &Point{
		Row: current.Row,
		Col: current.Col,
	}

	for ok := true; ok; ok = move(matrix, current) {
		matrix.setVisited(current.Row, current.Col, current.Direction)
	}

	return strconv.Itoa(matrix.getDistinctVisitedCount())
}

func (m *Matrix) at(row, col int) uint8 {
	if row < 0 || row >= m.Rows || col < 0 || col >= m.Cols {
		return 0
	}

	return m.Values[row][col]
}

func (m *Matrix) setVisited(row, col int, direction uint8) {
	m.Visited[(int(direction)*m.Rows*m.Cols)+(col*m.Rows)+row] = true
}

func (m *Matrix) getVisited(row, col int, direction uint8) bool {
	return m.Visited[(int(direction)*m.Rows*m.Cols)+(col*m.Rows)+row]
}

func move(m *Matrix, current *Point) bool {
	switch current.Direction {
	case DirectionNorth:
		if m.at(current.Row-1, current.Col) == Blocker {
			current.Direction = DirectionEast
		} else {
			current.Row--
		}
	case DirectionEast:
		if m.at(current.Row, current.Col+1) == Blocker {
			current.Direction = DirectionSouth
		} else {
			current.Col++
		}
	case DirectionSouth:
		if m.at(current.Row+1, current.Col) == Blocker {
			current.Direction = DirectionWest
		} else {
			current.Row++
		}
	case DirectionWest:
		if m.at(current.Row, current.Col-1) == Blocker {
			current.Direction = DirectionNorth
		} else {
			current.Col--
		}
	default:
		panic("unhandled default case")
	}

	return current.Row >= 0 && current.Row < m.Rows && current.Col >= 0 && current.Col < m.Cols
}

func (m *Matrix) getDistinctVisitedCount() int {
	count := 0

	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			if m.getVisited(row, col, DirectionNorth) || m.getVisited(row, col, DirectionEast) || m.getVisited(row, col, DirectionSouth) || m.getVisited(row, col, DirectionWest) {
				count++
			}
		}
	}

	return count
}

func Part02(_ []string) string {
	fmt.Printf("%v\n", start)
	return ""
}

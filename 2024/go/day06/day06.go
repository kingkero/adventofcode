package day06

import (
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
	DirectionNorth uint8 = 1 << iota
	DirectionEast
	DirectionSouth
	DirectionWest
)

const Blocker = 35 // string "#"
const Start = 94   // string "^"

func Part01(input []string) string {
	m := &Matrix{
		Rows: len(input),
		Cols: len(input[0]),
		Values: util.Filter(input, func(s string) bool {
			return s != ""
		}),
		Visited: make([]bool, len(input)*len(input[0])),
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

	for ok := true; ok; ok = move(m, current) {
		m.setVisited(current.Row, current.Col)
	}

	return strconv.Itoa(m.getVisitedCount())
}

func (m *Matrix) at(row, col int) uint8 {
	if row < 0 || row >= m.Rows || col < 0 || col >= m.Cols {
		return 0
	}

	return m.Values[row][col]
}

func (m *Matrix) setVisited(row, col int) {
	m.Visited[row*m.Cols+col] = true
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
	}

	return current.Row >= 0 && current.Row < m.Rows && current.Col >= 0 && current.Col < m.Cols
}

func (m *Matrix) getVisitedCount() int {
	count := 0
	for _, visited := range m.Visited {
		if visited {
			count++
		}
	}

	return count
}

func Part02(input []string) string {
	return ""
}

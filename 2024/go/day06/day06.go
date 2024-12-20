package day06

import (
	"strconv"
)

type Point struct {
	Row, Col  int
	Direction uint8
}

type Matrix struct {
	Rows, Cols int
	Values     []uint8
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
		Rows:    len(input),
		Cols:    len(input[0]),
		Values:  make([]uint8, len(input)*len(input[0])),
		Visited: make([]bool, len(input)*len(input[0])*int(AllDirections)),
	}

	current := &Point{Direction: DirectionNorth}
	for row, line := range input {
		for col, char := range line {
			matrix.Values[row*matrix.Cols+col] = uint8(char)

			if char == Start {
				current.Row = row
				current.Col = col
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

	return m.Values[row*m.Cols+col]
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
	// iterate over all visited except start
	// place block on them
	// try, on loops increment result

	localCopy := &Matrix{
		Rows:    matrix.Rows,
		Cols:    matrix.Cols,
		Values:  make([]uint8, len(matrix.Values)),
		Visited: make([]bool, len(matrix.Visited)),
	}

	result := 0

	var current *Point

	for row := 0; row < matrix.Rows; row++ {
		for col := 0; col < matrix.Cols; col++ {
			if row == start.Row && col == start.Col {
				continue
			}

			if matrix.getVisited(row, col, DirectionNorth) || matrix.getVisited(row, col, DirectionEast) || matrix.getVisited(row, col, DirectionSouth) || matrix.getVisited(row, col, DirectionWest) {
				localCopy.Values[row*matrix.Cols+col] = Blocker
				localCopy.Visited = make([]bool, matrix.Rows*matrix.Cols*int(AllDirections))

				current = &Point{
					Row:       start.Row,
					Col:       start.Col,
					Direction: start.Direction,
				}

				// fmt.Printf("%+v\n", localCopy)

			moving:
				for ok := true; ok; ok = move(localCopy, current) {
					if localCopy.getVisited(current.Row, current.Col, current.Direction) {
						result++

						break moving
					}

					localCopy.setVisited(current.Row, current.Col, current.Direction)
				}

				localCopy.Values[row*matrix.Cols+col] = matrix.at(row, col)
			}
		}
	}

	return strconv.Itoa(result)
}

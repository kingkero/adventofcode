package day04

import (
	"strconv"

	"github.com/kingkero/adventofcode/2024/go/util"
)

type Matrix struct {
	Rows, Cols     int
	Values         []string
	StartingPoints [][]int
}

func (m *Matrix) At(row, col int) string {
	return m.Values[row*m.Cols+col]
}

func (m *Matrix) Set(row, col int, val string) {
	m.Values[row*m.Cols+col] = val
}

func Part01(input []string) string {
	input = util.Filter(input, func(s string) bool {
		return s != ""
	})

	m := &Matrix{
		Rows:           len(input),
		Cols:           len(input[0]),
		Values:         make([]string, 0, len(input)*len(input[0])),
		StartingPoints: make([][]int, 0),
	}

	for row, line := range input {
		for col, char := range line {
			character := string(char)

			m.Values = append(m.Values, character)
			if character == "X" {
				m.StartingPoints = append(m.StartingPoints, []int{row, col})
			}
		}
	}

	result := 0
	for _, point := range m.StartingPoints {
		if isValidXmasToTheRight(m, point[0], point[1]) {
			result++
		}
		if isValidXmasToTheLeft(m, point[0], point[1]) {
			result++
		}
		if isValidXmasDown(m, point[0], point[1]) {
			result++
		}
		if isValidXmasUp(m, point[0], point[1]) {
			result++
		}
		if isValidXmasLeftDown(m, point[0], point[1]) {
			result++
		}
		if isValidXmasLeftUp(m, point[0], point[1]) {
			result++
		}
		if isValidXmasRightDown(m, point[0], point[1]) {
			result++
		}
		if isValidXmasRightUp(m, point[0], point[1]) {
			result++
		}
	}

	return strconv.Itoa(result)
}

func isValidXmasToTheRight(m *Matrix, row, col int) bool {
	if m.Cols <= col+3 {
		return false
	}

	return m.At(row, col+1)+m.At(row, col+2)+m.At(row, col+3) == "MAS"
}

func isValidXmasToTheLeft(m *Matrix, row, col int) bool {
	if col-3 < 0 {
		return false
	}

	return m.At(row, col-1)+m.At(row, col-2)+m.At(row, col-3) == "MAS"
}

func isValidXmasDown(m *Matrix, row, col int) bool {
	if m.Rows <= row+3 {
		return false
	}

	return m.At(row+1, col)+m.At(row+2, col)+m.At(row+3, col) == "MAS"
}

func isValidXmasUp(m *Matrix, row, col int) bool {
	if row-3 < 0 {
		return false
	}

	return m.At(row-1, col)+m.At(row-2, col)+m.At(row-3, col) == "MAS"
}

func isValidXmasRightDown(m *Matrix, row, col int) bool {
	if row+3 >= m.Rows || col+3 >= m.Cols {
		return false
	}

	return m.At(row+1, col+1)+m.At(row+2, col+2)+m.At(row+3, col+3) == "MAS"
}

func isValidXmasLeftUp(m *Matrix, row, col int) bool {
	if row-3 < 0 || col-3 < 0 {
		return false
	}

	return m.At(row-1, col-1)+m.At(row-2, col-2)+m.At(row-3, col-3) == "MAS"
}

func isValidXmasLeftDown(m *Matrix, row, col int) bool {
	if row+3 >= m.Rows || col-3 < 0 {
		return false
	}

	return m.At(row+1, col-1)+m.At(row+2, col-2)+m.At(row+3, col-3) == "MAS"
}

func isValidXmasRightUp(m *Matrix, row, col int) bool {
	if row-3 < 0 || col+3 >= m.Cols {
		return false
	}

	return m.At(row-1, col+1)+m.At(row-2, col+2)+m.At(row-3, col+3) == "MAS"
}

func Part02(input []string) string {
	input = util.Filter(input, func(s string) bool {
		return s != ""
	})

	m := &Matrix{
		Rows:           len(input),
		Cols:           len(input[0]),
		Values:         make([]string, 0, len(input)*len(input[0])),
		StartingPoints: make([][]int, 0),
	}

	for row, line := range input {
		for col, char := range line {
			character := string(char)

			m.Values = append(m.Values, character)
			if row > 0 && col > 0 && row < m.Rows-1 && col < m.Cols-1 && character == "A" {
				m.StartingPoints = append(m.StartingPoints, []int{row, col})
			}
		}
	}

	result := 0
	for _, point := range m.StartingPoints {
		switch m.At(point[0]-1, point[1]-1) + m.At(point[0]+1, point[1]+1) + m.At(point[0]+1, point[1]-1) + m.At(point[0]-1, point[1]+1) {
		case "MSMS":
			fallthrough
		case "MSSM":
			fallthrough
		case "SMMS":
			fallthrough
		case "SMSM":
			result++
		}
	}

	return strconv.Itoa(result)
}

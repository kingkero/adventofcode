package day16

import (
	"fmt"
	"log"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

// TODO: Extract Direction to util
type Direction int

const (
	North Direction = iota
	West
	South
	East
)

type Matrix struct {
	data    [][]string
	width   int
	height  int
	visited [][]int
}

func NewMatrix(lines []string) *Matrix {
	height, width := len(lines), len(lines[0])

	data := make([][]string, height)
	visited := make([][]int, height)

	for i, line := range lines {
		data[i] = strings.Split(line, "")
		visited[i] = make([]int, width)
	}

	return &Matrix{data, width, height, visited}
}

func (matrix *Matrix) walk(direction Direction, current util.Point) {
	if current.Row >= matrix.height || current.Col >= matrix.width || current.Row < 0 || current.Col < 0 {
		return
	}
	if matrix.visited[current.Row][current.Col] > 5 {
		/*
			if matrix.data[current.Row][current.Col] != "." {
				return
			}

			if matrix.visited[current.Row][current.Col] > 1 {
				return
			}
		*/
		return
	}

	matrix.visited[current.Row][current.Col]++

	switch matrix.data[current.Row][current.Col] {
	case ".":
		var next util.Point
		switch direction {
		case North:
			next = util.Point{Row: current.Row + 1, Col: current.Col}
			break
		case East:
			next = util.Point{Row: current.Row, Col: current.Col - 1}
			break
		case South:
			next = util.Point{Row: current.Row - 1, Col: current.Col}
			break
		case West:
			next = util.Point{Row: current.Row, Col: current.Col + 1}
			break
		}

		matrix.walk(direction, next)
		return

	case "/":
		switch direction {
		case North:
			matrix.walk(East, util.Point{Row: current.Row, Col: current.Col - 1})
			return
		case East:
			matrix.walk(North, util.Point{Row: current.Row + 1, Col: current.Col})
			return
		case South:
			matrix.walk(West, util.Point{Row: current.Row, Col: current.Col + 1})
			return
		case West:
			matrix.walk(South, util.Point{Row: current.Row - 1, Col: current.Col})
			return
		}
		break

	case "\\":
		switch direction {
		case North:
			matrix.walk(West, util.Point{Row: current.Row, Col: current.Col + 1})
			return
		case East:
			matrix.walk(South, util.Point{Row: current.Row - 1, Col: current.Col})
			return
		case South:
			matrix.walk(East, util.Point{Row: current.Row, Col: current.Col - 1})
			return
		case West:
			matrix.walk(North, util.Point{Row: current.Row + 1, Col: current.Col})
			return
		}
		break

	case "|":
		switch direction {
		case North:
			matrix.walk(North, util.Point{Row: current.Row + 1, Col: current.Col})
			return
		case East:
			matrix.walk(North, util.Point{Row: current.Row + 1, Col: current.Col})
			matrix.walk(South, util.Point{Row: current.Row - 1, Col: current.Col})
			return
		case South:
			matrix.walk(South, util.Point{Row: current.Row - 1, Col: current.Col})
			return
		case West:
			matrix.walk(North, util.Point{Row: current.Row + 1, Col: current.Col})
			matrix.walk(South, util.Point{Row: current.Row - 1, Col: current.Col})
			return
		}
		break

	case "-":
		switch direction {
		case North:
			matrix.walk(East, util.Point{Row: current.Row, Col: current.Col - 1})
			matrix.walk(West, util.Point{Row: current.Row, Col: current.Col + 1})
			return
		case East:
			matrix.walk(East, util.Point{Row: current.Row, Col: current.Col - 1})
			return
		case South:
			matrix.walk(East, util.Point{Row: current.Row, Col: current.Col - 1})
			matrix.walk(West, util.Point{Row: current.Row, Col: current.Col + 1})
			return
		case West:
			matrix.walk(West, util.Point{Row: current.Row, Col: current.Col + 1})
			return
		}
		break
	}

	return
}

func part01(matrix *Matrix) int {

	matrix.walk(West, util.Point{Row: 0, Col: 0})

	for _, row := range matrix.visited {
		for _, field := range row {
			fmt.Print(field)
		}
		fmt.Println()
	}

	return util.Sum(util.Map(matrix.visited, func(fields []int) int {
		return util.Sum(util.Map(fields, func(field int) int {
			if field > 0 {
				return 1
			}
			return 0
		}))
	}))
}

func part02(lines []string) int {
	return 0
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	matrix := NewMatrix(lines)

	return part01(matrix), part02(lines)
}

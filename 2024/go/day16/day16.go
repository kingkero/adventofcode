package day16

import (
	"fmt"

	"github.com/kingkero/adventofcode/2024/go/util"
)

type Field uint8

const (
	Fence Field = '#'
	Start Field = 'S'
	End   Field = 'E'
	Free  Field = '.'
)

type Matrix struct {
	width  int
	height int
	data   [][]Field
	start  util.Point
	end    util.Point
}

func Part01(input []string) string {
	matrix := &Matrix{
		width:  len(input[0]),
		height: len(input),
		start:  util.Point{},
		end:    util.Point{},
	}
	matrix.data = make([][]Field, matrix.width)

	for x, line := range input {
		matrix.data[x] = make([]Field, matrix.height)

		for y, char := range line {
			value := Field(char)
			matrix.data[x][y] = value

			if value == Start {
				matrix.start = util.Point{X: x, Y: y}
			} else if value == End {
				matrix.end = util.Point{X: x, Y: y}
			}
		}
	}

	fmt.Printf("Matrix: %v\n", matrix)

	return ""
}

func Part02(_ []string) string {
	return ""
}

package day13

import (
	"log"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

func NewMatrix(lines []string) *Matrix {
	height := len(lines)
	width := len(lines[0])

	data := make([][]string, height)
	for i, line := range lines {
		data[i] = strings.Split(line, "")
	}

	return &Matrix{data, width, height}
}

func ParseMatrixes(lines []string) []*Matrix {
	var matrixes []*Matrix

	startIndex := 0
	for i, line := range lines {
		if line == "" {
			matrixes = append(matrixes, NewMatrix(lines[startIndex:(i-startIndex)]))
			startIndex = i + 1
		}
	}
	matrixes = append(matrixes, NewMatrix(lines[startIndex:]))
	return matrixes
}

type Matrix struct {
	data   [][]string
	width  int
	height int
}

func (original Matrix) rotate() *Matrix {
	data := make([][]string, original.width)
	for row := 0; row < original.width; row++ {
		line := make([]string, original.height)
		for col := 0; col < original.height; col++ {
			line[col] = original.data[col][original.width-row-1]
		}
		data[row] = line
	}
	return &Matrix{data, original.height, original.width}
}

func part01(matrixes []*Matrix) int {
	return 0
}

func part02(lines []string) int {
	return 0
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	matrixes := ParseMatrixes(lines)

	return part01(matrixes), part02(lines)
}

package day13

import (
	"log"
	"math"
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

func canMirrorColAt(col int, row []string) bool {
	length := col + 1
	half := int(math.Floor(float64(len(row)) / 2.0))
	if col >= half {
		length = len(row) - 1 - col
	}

	for i := 0; i < length; i++ {
		if row[col-i] != row[col+i+1] {
			return false
		}
	}

	return true
}

func (matrix *Matrix) getLeftMirrorCol() int {
	// 0 cannot be a left mirror
	// width and width - 1 cannot be left mirrors
	possibles := make([]int, matrix.width-3)
	for i := range possibles {
		possibles[i] = i + 1
	}

	for _, row := range matrix.data {
		var newPossibles []int
		for _, possibleCol := range possibles {
			if canMirrorColAt(possibleCol, row) {
				newPossibles = append(newPossibles, possibleCol)
			}
		}

		if len(newPossibles) == 0 {
			return 0
		}
		possibles = newPossibles
	}

	if len(possibles) == 1 {
		return possibles[0]
	}

	return 0
}

func part01(matrixes []*Matrix) int {
	result := 0
	for _, matrix := range matrixes {
		if horizontalMirror := matrix.getLeftMirrorCol(); horizontalMirror > 0 {
			result += horizontalMirror + 1
			continue
		}

		if verticalMirror := matrix.rotate().getLeftMirrorCol(); verticalMirror > 0 {
			result += (verticalMirror + 1) * 100
		}
	}
	return result
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

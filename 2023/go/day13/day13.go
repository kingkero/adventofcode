package day13

import (
	"log"
	"math"
	"slices"
	"strconv"
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
			matrixes = append(matrixes, NewMatrix(lines[startIndex:i]))
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

func canMirrorColAt(col int, row []string, smudge bool) bool {
	length := col + 1
	half := int(math.Floor(float64(len(row)) / 2.0))
	var left []string

	if col < half {
		left = row[:length]
	} else {
		length = len(row) - 1 - col
		left = row[col-length+1 : col+1]
	}
	right := row[col+1 : col+1+length]

	rightCopy := make([]string, len(right))
	copy(rightCopy, right)
	slices.Reverse(rightCopy)

	leftString := strings.Join(left, "")
	rightString := strings.Join(rightCopy, "")

	if !smudge {
		return leftString == rightString
	}

	return util.Hamming(leftString, rightString) < 2
}

func getRowWithMissingReflectorCol(possiblesPerRow [][]int, col int) int {
	for i, possibles := range possiblesPerRow {
		if !slices.Contains(possibles, col) {
			return i
		}
	}
	return -1
}

func (matrix *Matrix) getLeftMirrorCol(smudge bool) int {
	// part 1:
	if !smudge {
		possibles := make([]int, matrix.width-1)
		for i := range possibles {
			possibles[i] = i
		}

		for _, row := range matrix.data {
			var newPossibles []int
			for _, possibleCol := range possibles {
				if canMirrorColAt(possibleCol, row, false) {
					newPossibles = append(newPossibles, possibleCol)
				}
			}

			if len(newPossibles) == 0 {
				return -1
			}

			possibles = newPossibles
		}

		if len(possibles) >= 1 {
			return possibles[0]
		}

		return -1
	}

	// get all possible reflections per row
	possiblesPerRow := make([][]int, len(matrix.data))
	for i, row := range matrix.data {
		var newPossibles []int
		for possibleCol := 0; possibleCol < matrix.width-1; possibleCol++ {
			if canMirrorColAt(possibleCol, row, false) {
				newPossibles = append(newPossibles, possibleCol)
			}
		}
		possiblesPerRow[i] = newPossibles
	}

	// get map col => amount of occurrences
	// if we see it matrix.height times, it is already a solution
	// if we see it matrix.height - 1 times, find the row that misses it and
	// 		check if we can reflect with smudge correction
	occurrences := util.SumOccurrences(possiblesPerRow)
	for col, occurres := range occurrences {
		if occurres == matrix.height-1 {
			row := getRowWithMissingReflectorCol(possiblesPerRow, col)
			if row > -1 {
				if canMirrorColAt(col, matrix.data[row], true) {
					return col
				}
			}
		}
	}

	return -1
}

func part01(matrixes []*Matrix) int {
	result := 0
	for i, matrix := range matrixes {
		if mirror := matrix.getLeftMirrorCol(false); mirror != -1 {
			result += mirror + 1
		} else if mirror := matrix.rotate().getLeftMirrorCol(false); mirror != -1 {
			result += (mirror + 1) * 100
		} else {
			panic("no reflection found for matrix " + strconv.Itoa(i))
		}
	}
	return result
}

func part02(matrixes []*Matrix) int {
	result := 0
	for i, matrix := range matrixes {
		if mirror := matrix.getLeftMirrorCol(true); mirror != -1 {
			result += mirror + 1
		} else if mirror := matrix.rotate().getLeftMirrorCol(true); mirror != -1 {
			result += (mirror + 1) * 100
		} else {
			panic("no reflection found for matrix " + strconv.Itoa(i))
		}
	}
	return result
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	matrixes := ParseMatrixes(lines)

	return part01(matrixes), part02(matrixes)
}

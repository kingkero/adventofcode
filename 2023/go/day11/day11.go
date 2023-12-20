package day11

import (
	"log"
	"math"
	"slices"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

func isRowEmpty(row []string) bool {
	for col := 0; col < len(row); col++ {
		if row[col] != "." {
			return false
		}
	}
	return true
}

func isColEmpty(fields [][]string, col int) bool {
	for row := 0; row < len(fields); row++ {
		if fields[row][col] != "." {
			return false
		}
	}
	return true
}

type Point struct {
	row int
	col int
}

type Pair struct {
	a *Point
	b *Point
}

type GalaxyImage struct {
	originalFields [][]string
	expandRows     []int
	expandCols     []int
	galaxies       []*Point
}

func NewGalaxyImage(lines []string) *GalaxyImage {
	var expandRows, expandCols []int
	var galaxies []*Point

	originalFields := make([][]string, len(lines))
	for row := 0; row < len(lines); row++ {
		originalFields[row] = strings.Split(lines[row], "")

		if isRowEmpty(originalFields[row]) {
			expandRows = append(expandRows, row+len(expandRows))
		}
	}

	for col := 0; col < len(originalFields[0]); col++ {
		if isColEmpty(originalFields, col) {
			expandCols = append(expandCols, col+len(expandCols))
		}
	}

	newRows := len(lines) + len(expandRows)
	newCols := len(lines[0]) + len(expandCols)

	referenceRow := 0
	for row := 0; row < newRows; row++ {
		referenceCol := 0
		for col := 0; col < newCols; col++ {
			if originalFields[referenceRow][referenceCol] == "#" {
				galaxies = append(galaxies, &Point{row, col})
			}
			referenceCol++

			if slices.Contains(expandCols, col) {
				col++
			}
		}
		referenceRow++

		if slices.Contains(expandRows, row) {
			row++
		}
	}

	return &GalaxyImage{originalFields, expandRows, expandCols, galaxies}
}

func (image GalaxyImage) getDistance(a, b *Point, factor int) int {
	diffRow := int(math.Abs(float64(image.galaxies[j].row - a.row)))
	diffCol := int(math.Abs(float64(image.galaxies[j].col - a.col)))

	if diffRow == 0 {
		return diffCol
	} else if diffCol == 0 {
		return diffRow
	} else {
		return diffCol + diffRow
	}
}

func (image GalaxyImage) getTotalDistances(factor int) int {
	distances := 0

	for i, a := range image.galaxies {
		for j := i + 1; j < len(image.galaxies); j++ {
			distances += image.getDistance(a, image.galaxies[j], factor)
		}
	}

	return distances
}

func part01(image *GalaxyImage) int {
	return image.getTotalDistances(1)
}

func part02(image *GalaxyImage) int {
	return image.getTotalDistances(10)
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	image := NewGalaxyImage(lines)

	return part01(image), part02(image)
}

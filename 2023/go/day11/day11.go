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
}

func NewGalaxyImage(lines []string) *GalaxyImage {
	var expandRows, expandCols []int

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

	return &GalaxyImage{originalFields, expandRows, expandCols}
}

func (image *GalaxyImage) getDistance(a, b *Point) int {
	diffRow := int(math.Abs(float64(b.row - a.row)))
	diffCol := int(math.Abs(float64(b.col - a.col)))

	if diffRow == 0 {
		return diffCol
	}
	if diffCol == 0 {
		return diffRow
	}
	return diffCol + diffRow
}

func (image *GalaxyImage) getTotalGalaxiesDistance(expand int) int {
	newRows := len(image.originalFields) + len(image.expandRows)
	newCols := len(image.originalFields[0]) + len(image.expandCols)
	var galaxies []*Point

	referenceRow := 0
	for row := 0; row < newRows; row++ {
		referenceCol := 0
		for col := 0; col < newCols; col++ {
			if image.originalFields[referenceRow][referenceCol] == "#" {
				galaxies = append(galaxies, &Point{row, col})
			}
			referenceCol++

			if slices.Contains(image.expandCols, col) {
				col++
			}
		}
		referenceRow++

		if slices.Contains(image.expandRows, row) {
			row++
		}
	}

	distances := 0

	for i, a := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			distances += image.getDistance(a, galaxies[j])
		}
	}

	return distances
}

func part01(image *GalaxyImage) int {
	return image.getTotalGalaxiesDistance(1)
}

func part02(image *GalaxyImage) int {
	return image.getTotalGalaxiesDistance(9)
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	image := NewGalaxyImage(lines)

	return part01(image), part02(image)
}

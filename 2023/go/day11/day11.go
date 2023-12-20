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
	galaxies       []*Point
	emptyRows      []int
	emptyCols      []int
}

func NewGalaxyImage(lines []string) *GalaxyImage {
	var galaxies []*Point
	var emptyRows, emptyCols []int

	originalFields := make([][]string, len(lines))
	for row := 0; row < len(lines); row++ {
		originalFields[row] = strings.Split(lines[row], "")
		if isRowEmpty(originalFields[row]) {
			emptyRows = append(emptyRows, row)
		}
	}

	for row := 0; row < len(originalFields); row++ {
		for col := 0; col < len(originalFields[0]); col++ {
			if row == 0 {
				if isColEmpty(originalFields, col) {
					emptyCols = append(emptyCols, col)
				}
			}
			if originalFields[row][col] == "#" {
				galaxies = append(galaxies, &Point{row, col})
			}
		}
	}

	return &GalaxyImage{originalFields, galaxies, emptyRows, emptyCols}
}

func (image GalaxyImage) getDistance(a, b *Point, factor int) int {
	startRow := int(math.Min(float64(a.row), float64(b.row)))
	diffRow := int(math.Abs(float64(b.row - a.row)))
	maxDiffRow := diffRow

	startCol := int(math.Min(float64(a.col), float64(b.col)))
	diffCol := int(math.Abs(float64(b.col - a.col)))
	maxDiffCol := diffCol

	for i := 1; i < maxDiffRow; i++ {
		if slices.Contains(image.emptyRows, startRow+i) {
			diffRow += factor
		}
	}

	for i := 1; i < maxDiffCol; i++ {
		if slices.Contains(image.emptyCols, startCol+i) {
			diffCol += factor
		}
	}

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

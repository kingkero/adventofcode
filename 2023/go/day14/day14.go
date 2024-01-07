package day14

import (
	"log"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

type Platform struct {
	matrix [][]string
	cols   [][]string
}

func NewPlatform(lines []string) *Platform {
	matrix := make([][]string, len(lines))

	for i, line := range lines {
		matrix[i] = strings.Split(line, "")
	}

	return &Platform{matrix, nil}
}

func (platform Platform) getColumns() [][]string {
	cols := make([][]string, len(platform.matrix[0]))

	for col := 0; col < len(cols); col++ {
		newCol := make([]string, len(platform.matrix))
		for row := 0; row < len(newCol); row++ {
			newCol[row] = platform.matrix[row][col]
		}
		cols[col] = newCol
	}

	return cols
}

func getNextRoundedRockIndex(col []string, startIndex int) int {
	for i := startIndex; i < len(col); i++ {
		if col[i] == "#" {
			return -1
		}
		if col[i] == "O" {
			return i
		}
	}
	return -1
}

func (platform Platform) tiltNorth() *Platform {
	cols := platform.getColumns()

	// similar to bubblesort, move "O" up if encoutering "."
	for _, col := range cols {
		for j := 0; j < len(col)-1; j++ {
			if col[j] == "O" || col[j] == "#" {
				continue
			}

			if nextRoundRock := getNextRoundedRockIndex(col, j); nextRoundRock > -1 {
				col[j] = "O"
				col[nextRoundRock] = "."
			}
		}
	}

	return &Platform{nil, cols}
}

func (platform Platform) getTotalLoad() int {
	result := 0
	height := len(platform.cols[0])

	for _, col := range platform.cols {
		for i, element := range col {
			if element == "O" {
				result += height - i
			}
		}
	}

	return result
}

func part01(platform Platform) int {
	return platform.tiltNorth().getTotalLoad()
}

func part02(platform Platform) int {
	return 0
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	platform := NewPlatform(lines)

	// dump.P(platform)

	return part01(*platform), part02(*platform)
}

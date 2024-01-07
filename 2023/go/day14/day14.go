package day14

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

type Platform struct {
	matrix [][]string
}

func NewPlatform(lines []string) *Platform {
	matrix := make([][]string, len(lines))

	for i, line := range lines {
		matrix[i] = strings.Split(line, "")
	}

	return &Platform{matrix}
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

type Direction int

const (
	North Direction = iota
	West
	South
	East
)

func (platform *Platform) tilt(direction Direction) {
	data := platform.matrix

	if direction == North {
		data = platform.getColumns()
	}
	if direction == South {
		data = platform.getColumns()

		for _, col := range data {
			slices.Reverse(col)
		}
	}

	for _, list := range data {
		for j := 0; j < len(list)-1; j++ {
			if list[j] == "O" || list[j] == "#" {
				continue
			}

			if found := getNextRoundedRockIndex(list, j); found != -1 {
				list[j] = "O"
				list[found] = "."
			}
		}
	}

	if direction == North || direction == South {
		for i, col := range data {
			if direction == South {
				slices.Reverse(col)
			}

			for j, element := range col {
				platform.matrix[j][i] = element
			}
		}
	}
}

func (platform Platform) getTotalLoad() int {
	result := 0
	height := len(platform.matrix)

	for i, row := range platform.matrix {
		for _, element := range row {
			if element == "O" {
				result += height - i
			}
		}
	}

	return result
}

func (platform *Platform) cycle() {
	platform.tilt(North)
	for _, row := range platform.matrix {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
	platform.tilt(West)
	for _, row := range platform.matrix {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
	platform.tilt(South)
	for _, row := range platform.matrix {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
	// platform.tiltSouth()
	// fmt.Println()
	// platform.tiltEast()
	// fmt.Println()
}

func part01(platform Platform) int {
	// platform.tilt(North)
	return platform.getTotalLoad()
}

func part02(platform Platform) int {
	for i := 0; i < 1; i++ {
		platform.cycle()
	}
	return platform.getTotalLoad()
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

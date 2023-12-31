package day14

import (
	"log"
	"strings"

	"github.com/gookit/goutil/dump"
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

func part01(platform Platform) int {
	return 0
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

	dump.P(platform)

	return part01(*platform), part02(*platform)
}

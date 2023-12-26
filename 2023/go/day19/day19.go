package day19

import (
	"log"

	"github.com/kingkero/adventofcode/2023/go/util"
)

func part01(lines []string) int {
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

	return part01(lines), part02(lines)
}

package day05

import (
	"log"
	"strings"

	"github.com/gookit/goutil/dump"
	"github.com/kingkero/adventofcode/2023/go/util"
)

func part01(lines []string) int {
	result := 0

	seeds := strings.Split(strings.Split(lines[0], ": ")[1], " ")

	seedsToSoil := make([][]int, 2)
	seedsToSoil[0] = util.Map(strings.Split(lines[3], " "), util.ParseInt)
	seedsToSoil[1] = util.Map(strings.Split(lines[4], " "), util.ParseInt)

	dump.P(seeds)
	dump.P(seedsToSoil)

	return result
}

func part02(lines []string) int {
	result := 0

	return result
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	return part01(lines), part02(lines)
}

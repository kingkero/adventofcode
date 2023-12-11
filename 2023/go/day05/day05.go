package day05

import (
	"log"
	"strings"

	"github.com/gookit/goutil/dump"
	"github.com/kingkero/adventofcode/2023/go/util"
)

func getMappedId(originalId int, data [][]int) int {
	result := originalId

	for _, rules := range data {
		if originalId < rules[1] {
			continue
		}
		if rules[1]+rules[2]-1 < originalId {
			continue
		}

		return originalId - rules[1] + rules[0]
	}

	return result
}

func part01(lines []string) int {
	result := 0

	seeds := util.Map(strings.Split(strings.Split(lines[0], ": ")[1], " "), util.ParseInt)

	seedsToSoilData := make([][]int, 2)
	seedsToSoilData[0] = util.Map(strings.Split(lines[3], " "), util.ParseInt)
	seedsToSoilData[1] = util.Map(strings.Split(lines[4], " "), util.ParseInt)

	// dump.P(seeds)
	// dump.P(seedsToSoilData)
	dump.P(getMappedId(seeds[0], seedsToSoilData))
	dump.P(getMappedId(seeds[1], seedsToSoilData))
	dump.P(getMappedId(seeds[2], seedsToSoilData))
	dump.P(getMappedId(seeds[3], seedsToSoilData))

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

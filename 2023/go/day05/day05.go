package day05

import (
	"log"
	"slices"
	"strings"

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

func getMapData(start int, lines []string) ([][]int, int) {
	var data [][]int
	for i := start; i < len(lines); i++ {
		// fmt.Println(lines[i])
		if lines[i] == "" {
			return data, i
		}

		data = append(data, util.Map(strings.Split(lines[i], " "), util.ParseInt))
	}

	return data, -1
}

func getLocationIds(lines []string) []int {
	ids := util.Map(strings.Split(strings.Split(lines[0], ": ")[1], " "), util.ParseInt)

	var data [][]int
	lastLine := 1

	for {
		data, lastLine = getMapData(lastLine+2, lines)
		tmp := make([]int, len(ids))
		for i, id := range ids {
			tmp[i] = getMappedId(id, data)
		}

		ids = tmp

		if lastLine == -1 {
			return ids
		}
	}
}

func part01(lines []string) int {
	locationIds := getLocationIds(lines)

	return slices.Min(locationIds)
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

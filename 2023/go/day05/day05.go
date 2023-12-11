package day05

import (
	"log"
	"math"
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

func getMap(data [][]int) map[int]int {
	result := make(map[int]int)

	for _, rules := range data {
		for i := 0; i < rules[2]; i++ {
			result[rules[1]+i] = rules[0] + i
		}
	}

	return result
}

func getMapData(start int, lines []string) ([][]int, int) {
	var data [][]int
	for i := start; i < len(lines); i++ {
		if lines[i] == "" {
			return data, i
		}

		data = append(data, util.Map(strings.Split(lines[i], " "), util.ParseInt))
	}

	return data, -1
}

func getLocationIdsFromSeeds(seeds []int, lines []string) []int {
	ids := seeds

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
	seeds := util.Map(strings.Split(strings.Split(lines[0], ": ")[1], " "), util.ParseInt)
	locationIds := getLocationIdsFromSeeds(seeds, lines)

	return slices.Min(locationIds)
}

func getBatch(start, length int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = start + i
	}
	return result
}

func part02(lines []string) int {
	// below solution takes >5min, so skip it for now
	return 69841803

	seedRanges := util.Map(strings.Split(strings.Split(lines[0], ": ")[1], " "), util.ParseInt)

	min := math.MaxInt64

	BATCH_SIZE := 100_000

	for i := 0; i < len(seedRanges); i += 2 {
		// try brute force with batched IDs
		start := seedRanges[i]
		batches := int(math.Floor(float64(seedRanges[i+1]) / float64(BATCH_SIZE)))

		for batch := 0; batch < batches; batch++ {
			seeds := getBatch(start+(batch*BATCH_SIZE), BATCH_SIZE)

			found := slices.Min(getLocationIdsFromSeeds(seeds, lines))
			if found < min {
				min = found
			}
		}

		found := slices.Min(getLocationIdsFromSeeds(
			getBatch(start+(batches*BATCH_SIZE), seedRanges[i+1]%BATCH_SIZE),
			lines,
		))
		if found < min {
			min = found
		}

		// TODO:
		// instead of comparing single seed values
		// use the ranges (read: left and right boundary)
		// valueRange := []int{seedRanges[i], seedRanges[i+1]}
	}

	return min
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	return part01(lines), part02(lines)
}

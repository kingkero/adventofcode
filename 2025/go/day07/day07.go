package day07

import (
	"strconv"
	"strings"
)

func Part01(input []string) string {
	positionToBeam := make(map[int]bool)
	width := len(input[0])
	splits := 0

	for i, element := range input[0] {
		if element == 'S' {
			positionToBeam[i] = true
			break
		}
	}

	for _, line := range input[1:] {
		if strings.Trim(line, ".") == "" {
			continue
		}

		for i := 0; i < width; i++ {
			if _, ok := positionToBeam[i]; ok {
				if line[i] == '^' {
					positionToBeam[i-1] = true
					positionToBeam[i+1] = true
					delete(positionToBeam, i)
					splits++
				}
			}
		}
	}

	return strconv.Itoa(splits)
}

type cacheKey struct {
	beamPosition int
	rowIndex     int
}

func getTimelines(beamPosition int, rowIndex int, input []string, cache map[cacheKey]int) int {
	key := cacheKey{beamPosition, rowIndex}
	if val, ok := cache[key]; ok {
		return val
	}

	for i := rowIndex + 1; i < len(input); i++ {
		line := input[i]
		if strings.Trim(line, ".") == "" {
			continue
		}

		if line[beamPosition] != '^' {
			continue
		}

		result := getTimelines(beamPosition-1, i, input, cache) + getTimelines(beamPosition+1, i, input, cache)
		cache[key] = result
		return result
	}

	cache[key] = 1
	return 1
}

func Part02(input []string) string {
	beamPosition := 0

	for i, element := range input[0] {
		if element == 'S' {
			beamPosition = i
			break
		}
	}

	cache := make(map[cacheKey]int)
	return strconv.Itoa(getTimelines(beamPosition, 0, input, cache))
}

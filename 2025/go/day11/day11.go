package day11

import (
	"strconv"
	"strings"
)

func findPaths(devices map[string][]string, current string, cache map[string]int) int {
	if value, ok := cache[current]; ok {
		return value
	}

	if len(devices[current]) == 0 {
		return 0
	}

	result := 0
	for _, device := range devices[current] {
		if device == "out" {
			result++
		} else {
			result += findPaths(devices, device, cache)
		}
	}

	cache[current] = result
	return result
}

func Part01(input []string) string {
	cache := make(map[string]int)
	devices := make(map[string][]string)

	for _, line := range input {
		parts := strings.Split(line, " ")
		devices[parts[0][:len(parts[0])-1]] = parts[1:]
	}

	return strconv.Itoa(findPaths(devices, "you", cache))
}

func findPathsPart02(devices map[string][]string, current string, visitedDac, visitedFft bool, cache map[string]int) int {
	if current == "dac" {
		visitedDac = true
	}
	if current == "fft" {
		visitedFft = true
	}

	cacheKey := current
	if visitedDac {
		cacheKey += "|d"
	}
	if visitedFft {
		cacheKey += "|f"
	}

	if value, ok := cache[cacheKey]; ok {
		return value
	}

	if len(devices[current]) == 0 {
		return 0
	}

	result := 0
	for _, device := range devices[current] {
		if device == "out" {
			if visitedDac && visitedFft {
				result++
			}
		} else {
			result += findPathsPart02(devices, device, visitedDac, visitedFft, cache)
		}
	}

	cache[cacheKey] = result
	return result
}

func Part02(input []string) string {
	cache := make(map[string]int)
	devices := make(map[string][]string)

	for _, line := range input {
		parts := strings.Split(line, " ")
		devices[parts[0][:len(parts[0])-1]] = parts[1:]
	}

	return strconv.Itoa(findPathsPart02(devices, "svr", false, false, cache))
}

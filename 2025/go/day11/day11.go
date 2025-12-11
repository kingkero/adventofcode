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

func Part02(input []string) string {
	return strconv.Itoa(0)
}

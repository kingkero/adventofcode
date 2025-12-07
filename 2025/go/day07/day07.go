package day07

import (
	"strconv"
)

func Part01(input []string) string {
	positionToBeam := make(map[int]bool)
	width := len(input[0])
	splits := 0

	for i, element := range input[0] {
		if element == 'S' {
			positionToBeam[i] = true
		}
	}

	for _, line := range input[1:] {
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

func Part02(input []string) string {
	return strconv.Itoa(0)
}

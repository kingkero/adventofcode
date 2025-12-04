package day04

import (
	"strconv"
)

type position struct {
	col, row int
}

func Part01(input []string) string {
	// width := len(input[0])
	// height := len(input)

	sum := 0

	paperRollsToAdjacents := make(map[position]int)

	for row, line := range input {
		for col, char := range line {
			if char != '@' {
				continue
			}

			paperRollsToAdjacents[position{col, row}] = 0
		}
	}

	diffsToCheck := [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{+1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	for pos, _ := range paperRollsToAdjacents {
		for _, diff := range diffsToCheck {
			if _, found := paperRollsToAdjacents[position{pos.col - diff[0], pos.row - diff[1]}]; found {
				paperRollsToAdjacents[position{pos.col - diff[0], pos.row - diff[1]}]++
			}
		}
	}

	for _, adjacents := range paperRollsToAdjacents {
		if adjacents < 4 {
			sum++
		}
	}

	return strconv.Itoa(sum)
}

func Part02(input []string) string {
	return ""
}

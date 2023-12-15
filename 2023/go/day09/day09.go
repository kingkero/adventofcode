package day09

import (
	"log"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

// returns (nextGuess, isOnlyZeroes)
func guessNextValue(series []int) int {
	var guessLines [][]int
	current := series

	guessLines = append(guessLines, series)
	for {
		next, isAllZeroes := getDifference(current)
		guessLines = append(guessLines, next)
		current = next
		if isAllZeroes {
			break
		}
	}

	nextVal := 0
	for i := len(guessLines) - 1; i > 0; i-- {
		nextVal = guessLines[i-1][len(guessLines[i-1])-1] + guessLines[i][len(guessLines[i])-1]
		guessLines[i-1] = append(guessLines[i-1], nextVal)
	}

	return nextVal
}

// Get difference between elements
func getDifference(series []int) ([]int, bool) {
	nextGuess := make([]int, len(series)-1)
	allZeroes := true
	for i := 0; i < len(series)-1; i++ {
		nextGuess[i] = series[i+1] - series[i]
		if allZeroes && nextGuess[i] != 0 {
			allZeroes = false
		}
	}
	return nextGuess, allZeroes
}

func part01(lines []string) int {
	nextValues := make([]int, len(lines))
	for i, line := range lines {
		nextValues[i] = guessNextValue(util.Map(strings.Split(line, " "), util.ParseInt))
	}

	result := 0
	for _, val := range nextValues {
		result += val
	}

	return result
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

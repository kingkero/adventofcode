package day09

import (
	"log"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

func getGuessLines(series []int) [][]int {
	var result [][]int
	current := series

	result = append(result, series)
	for {
		next, isAllZeroes := getDifference(current)
		result = append(result, next)
		current = next
		if isAllZeroes {
			break
		}
	}

	return result
}

func guessNextValue(series []int) int {
	guessLines := getGuessLines(series)

	nextVal := 0
	for i := len(guessLines) - 1; i > 0; i-- {
		nextVal = guessLines[i-1][len(guessLines[i-1])-1] + guessLines[i][len(guessLines[i])-1]
		guessLines[i-1] = append(guessLines[i-1], nextVal)
	}

	return nextVal
}

// Difference between each value in the slice, second return argument is true if
// all differences are 0s.
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
	result := 0
	for _, line := range lines {
		result += guessNextValue(util.Map(strings.Split(line, " "), util.ParseInt))
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

package day06

import (
	"slices"
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2025/go/util"
)

func Part01(input []string) string {
	operators := slices.DeleteFunc(strings.Split(input[len(input)-1], " "), func(s string) bool {
		return s == ""
	})

	problemResults := util.Map(slices.DeleteFunc(strings.Split(input[0], " "), func(s string) bool {
		return s == ""
	}), util.ParseInt)

	for _, line := range input[1 : len(input)-1] {
		parts := util.Map(slices.DeleteFunc(strings.Split(line, " "), func(s string) bool {
			return s == ""
		}), util.ParseInt)

		for i, operator := range operators {
			if operator == "+" {
				problemResults[i] += parts[i]
			} else {
				problemResults[i] *= parts[i]
			}
		}
	}

	sum := 0
	for _, result := range problemResults {
		sum += result
	}

	return strconv.Itoa(sum)
}

func Part02(input []string) string {
	operators := slices.DeleteFunc(strings.Split(input[len(input)-1], " "), func(s string) bool {
		return s == ""
	})

	sum := 0
	currentOperatorIndex := len(operators) - 1
	current := -1

	for col := len(input[0]) - 1; col >= 0; col-- {
		var parts strings.Builder
		for row := 0; row < len(input)-1; row++ {
			parts.WriteByte(input[row][col])
		}

		concatenated := strings.Trim(parts.String(), " ")

		if current == -1 {
			current = util.ParseInt(concatenated)
			continue
		}

		if concatenated == "" {
			sum += current
			current = -1
			currentOperatorIndex--
			continue
		}

		if operators[currentOperatorIndex] == "+" {
			current += util.ParseInt(concatenated)
		} else {
			current *= util.ParseInt(concatenated)
		}
	}

	sum += current

	return strconv.Itoa(sum)
}

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
	return strconv.Itoa(0)
}

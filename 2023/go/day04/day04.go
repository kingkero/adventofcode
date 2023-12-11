package day04

import (
	"log"
	"math"
	"regexp"
	"slices"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

func getScratchCardNumbers(line string) ([]string, []string) {
	numbersRegexp := regexp.MustCompile("\\d+")
	data := strings.Split(strings.Split(line, ":")[1], "|")
	return numbersRegexp.FindAllString(data[0], -1), numbersRegexp.FindAllString(data[1], -1)
}

func part01(lines []string) int {
	result := 0

	for _, line := range lines {
		winners, mine := getScratchCardNumbers(line)
		wins := 0
		for _, number := range mine {
			if slices.Contains(winners, number) {
				wins++
			}
		}
		if wins > 0 {
			result += int(math.Pow(2.0, float64(wins-1)))
		}
	}

	return result
}

func part02(lines []string) int {
	result := 0

	copies := make([]int, len(lines))

	for cardIndex, line := range lines {
		copies[cardIndex]++

		winners, mine := getScratchCardNumbers(line)
		wins := 0
		for _, number := range mine {
			if slices.Contains(winners, number) {
				wins++
				copies[cardIndex+wins] += copies[cardIndex]
			}
		}

		result += copies[cardIndex]
	}

	return result
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	return part01(lines), part02(lines)
}

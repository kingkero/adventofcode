package day06

import (
	"log"
	"regexp"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

func getWinningCharges(time, distance int) []int {
	var result []int

	for hold := 1; hold < time; hold++ {
		run := hold * (time - hold)
		if run > distance {
			result = append(result, hold)
		}
	}

	return result
}

func part01(lines []string) int {
	numbersRegexp := regexp.MustCompile("\\d+")

	times := util.Map(numbersRegexp.FindAllString(strings.Split(lines[0], ": ")[1], -1), util.ParseInt)
	distances := util.Map(numbersRegexp.FindAllString(strings.Split(lines[1], ": ")[1], -1), util.ParseInt)

	result := 1

	for i := range times {
		result *= len(getWinningCharges(times[i], distances[i]))
	}

	return result
}

func part02(lines []string) int {
	whitespaceRegexp := regexp.MustCompile("\\s+")
	time := util.ParseInt(whitespaceRegexp.ReplaceAllString(strings.Split(lines[0], ": ")[1], ""))
	distance := util.ParseInt(whitespaceRegexp.ReplaceAllString(strings.Split(lines[1], ": ")[1], ""))

	return len(getWinningCharges(time, distance))
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	return part01(lines), part02(lines)
}

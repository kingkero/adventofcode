package day07

import (
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2024/go/util"
)

const (
	OperatorAdd uint8 = iota
	OperatorMultiply

	AllOperators
)

type Calibration struct {
	result    int
	parts     []int
	operators []uint8
}

func Part01(input []string) string {
	result := 0

	for _, line := range input {
		resultAndParts := strings.Split(line, ": ")

		calibration := Calibration{
			result: util.ParseInt(resultAndParts[0]),
			parts:  util.Map(strings.Split(resultAndParts[1], " "), util.ParseInt),
		}

		if calibration.isSolvable() {
			result += calibration.result
		}
	}

	return strconv.Itoa(result)
}

func (calibration *Calibration) isSolvable() bool {
	return false
}

func Part02(_ []string) string {
	return ""
}

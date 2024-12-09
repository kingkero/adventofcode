package day07

import (
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2024/go/util"
)

const (
	OperatorAdd uint8 = iota
	OperatorMultiply
	OperatorAppend

	AllOperators
)

type Calibration struct {
	result int
	parts  []int
}

func Part01(input []string) string {
	result := 0

	for _, line := range input {
		resultAndParts := strings.Split(line, ": ")

		calibration := &Calibration{
			result: util.ParseInt(resultAndParts[0]),
			parts:  util.Map(strings.Split(resultAndParts[1], " "), util.ParseInt),
		}

		if calibration.isSolvable(false) {
			result += calibration.result
		}
	}

	return strconv.Itoa(result)
}

func (calibration *Calibration) isSolvable(withAppend bool) bool {
	for _, operators := range generateOperators(withAppend, len(calibration.parts)-1) {
		tmpResult := calibration.parts[0]

		for i := 1; i < len(calibration.parts); i++ {
			switch operators[i-1] {
			case OperatorAdd:
				tmpResult = tmpResult + calibration.parts[i]
			case OperatorMultiply:
				tmpResult = tmpResult * calibration.parts[i]
			case OperatorAppend:
				tmpResult = util.ParseInt(strconv.Itoa(tmpResult) + strconv.Itoa(calibration.parts[i]))
			}

		}

		if tmpResult == calibration.result {
			return true
		}
	}

	return false
}

func generateOperators(withAppend bool, length int) [][]uint8 {
	if length == 1 {
		if withAppend {
			return [][]uint8{
				{OperatorAdd},
				{OperatorMultiply},
				{OperatorAppend},
			}
		}

		return [][]uint8{
			{OperatorAdd},
			{OperatorMultiply},
		}
	}

	previousResult := generateOperators(withAppend, length-1)
	result := make([][]uint8, 0, len(previousResult)*int(AllOperators))

	for _, prev := range previousResult {
		result = append(result, append([]uint8{OperatorAdd}, prev...))
		result = append(result, append([]uint8{OperatorMultiply}, prev...))

		if withAppend {
			result = append(result, append([]uint8{OperatorAppend}, prev...))
		}
	}

	return result
}

func Part02(input []string) string {
	result := 0

	for _, line := range input {
		resultAndParts := strings.Split(line, ": ")

		calibration := &Calibration{
			result: util.ParseInt(resultAndParts[0]),
			parts:  util.Map(strings.Split(resultAndParts[1], " "), util.ParseInt),
		}

		if calibration.isSolvable(true) {
			result += calibration.result
		}
	}

	return strconv.Itoa(result)
}

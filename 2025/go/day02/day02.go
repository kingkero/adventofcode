package day02

import (
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2025/go/util"
)

func patternRepeats(pattern string, rest string) bool {
	patternLength := len(pattern)
	restLength := len(rest)

	if restLength%patternLength != 0 {
		return false
	}

	for i := 0; i < restLength/patternLength; i++ {
		for j := 0; j < patternLength; j++ {
			if rest[i+j:i+j+1] != pattern[j:j+1] {
				return false
			}
		}
	}

	return true
}

func isInvalidID(id string) bool {
	length := len(id)

	// required?
	if length <= 1 {
		return true
	}

	for i := 1; i < length; i++ {
		if patternRepeats(id[0:i], id[i:]) {
			return true
		}
	}

	return false
}

func Part01(input []string) string {
	inputRanges := strings.Split(input[0], ",")

	sumOfInvalids := 0

	for _, line := range inputRanges {
		parts := strings.Split(line, "-")

		start := util.ParseInt(parts[0])
		end := util.ParseInt(parts[1])

		for i := start; i <= end; i++ {
			if isInvalidID(strconv.Itoa(i)) {
				sumOfInvalids += i
			}
		}
	}

	return strconv.Itoa(sumOfInvalids)
}

func Part02(input []string) string {
	return ""
}

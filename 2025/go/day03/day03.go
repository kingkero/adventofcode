package day03

import (
	"strconv"

	"github.com/kingkero/adventofcode/2025/go/util"
)

// returns (value, position)
func findLargest(input string) (int, int) {
	length := len(input)
	if length == 1 {
		return util.ParseInt(input), 0
	}

	largest := util.ParseInt(input[0:1])
	largestPos := 0

	for i := 1; i < length; i++ {
		element := util.ParseInt(input[i : i+1])
		if element > largest {
			largest = element
			largestPos = i
		}
	}

	return largest, largestPos
}

func Part01(input []string) string {
	sum := 0

	for _, line := range input {
		firstLargest, pos := findLargest(line[:len(line)-1])
		secondLargest, _ := findLargest(line[pos+1:])

		jolt := firstLargest*10 + secondLargest
		sum += jolt
	}
	return strconv.Itoa(sum)
}

func Part02(input []string) string {
	return ""
}

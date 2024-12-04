package day03

import (
	"regexp"
	"strconv"

	"github.com/kingkero/adventofcode/2024/go/util"
)

func Part01(input []string) string {
	result := 0
	r, _ := regexp.Compile("mul\\((\\d+),(\\d+)\\)")

	for _, line := range input {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			result += util.ParseInt(match[1]) * util.ParseInt(match[2])
		}
	}

	return strconv.Itoa(result)
}

func Part02(input []string) string {
	return ""
}

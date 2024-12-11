package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2024/go/util"
)

func Part01(input []string) string {
	rules := make(map[uint8][]uint8)

	var left uint8
	var right uint8

	var lastRuleKey int

	for key, line := range input {
		if line == "" {
			lastRuleKey = key
			break
		}

		fmt.Sscanf(line, "%d|%d", &left, &right)
		if _, ok := rules[left]; !ok {
			rules[left] = make([]uint8, 0)
		}

		rules[left] = append(rules[left], right)
	}

	result := 0

	for key := lastRuleKey + 1; key < len(input); key++ {
		parts := util.Map(strings.Split(input[key], ","), util.ParseUint8)

		if isValidOrder(rules, parts) {
			result += int(parts[len(parts)/2])
		}
	}

	return strconv.Itoa(result)
}

func isValidOrder(rules map[uint8][]uint8, parts []uint8) bool {
	for i := 1; i < len(parts); i++ {
		values, ok := rules[parts[i]]

		if !ok {
			continue
		}

		for j := i - 1; j >= 0; j-- {
			if slices.Contains(values, parts[j]) {
				return false
			}
		}
	}

	return true
}

func Part02(input []string) string {
	return ""
}

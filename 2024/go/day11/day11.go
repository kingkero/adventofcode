package day11

import (
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2024/go/util"
)

func Part01(input []string) string {
	stones := util.Map(strings.Split(input[0], " "), util.ParseUint64)
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	return strconv.Itoa(len(stones))
}

func blink(stones []uint64) []uint64 {
	result := make([]uint64, 0, len(stones))

	for _, stone := range stones {
		if stone == 0 {
			result = append(result, 1)
			continue
		}

		val := strconv.Itoa(int(stone))
		if len(val)%2 == 0 {
			result = append(result, util.ParseUint64(val[0:len(val)/2]))
			result = append(result, util.ParseUint64(val[len(val)/2:]))

			continue
		}

		result = append(result, stone*2024)
	}

	return result
}

func Part02(_ []string) string {
	return ""
}

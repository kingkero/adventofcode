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

func smartBlink(stones []uint64, amountOfBlinks int) uint64 {
	result := uint64(0)

	for _, stone := range stones {
		result += smartBlinkStone(stone, amountOfBlinks)
	}

	return result
}

var alreadyBlinked = make(map[uint64]map[int]uint64)

func smartBlinkStone(stone uint64, amountOfBlinks int) uint64 {
	if amountOfBlinks == 0 {
		return 1
	}

	if _, ok := alreadyBlinked[stone]; !ok {
		alreadyBlinked[stone] = make(map[int]uint64)
	}

	if val, ok := alreadyBlinked[stone][amountOfBlinks]; ok {
		return val
	}

	a, b, split := applyRule(stone)

	if split {
		return smartBlinkStone(a, amountOfBlinks-1) + smartBlinkStone(b, amountOfBlinks-1)
	}

	result := smartBlinkStone(a, amountOfBlinks-1)

	if _, ok := alreadyBlinked[stone][amountOfBlinks]; !ok {
		alreadyBlinked[stone][amountOfBlinks] = result
	}

	return result
}

func applyRule(stone uint64) (uint64, uint64, bool) {
	if stone == 0 {
		return 1, 0, false
	}

	val := strconv.Itoa(int(stone))
	if len(val)%2 == 0 {
		return util.ParseUint64(val[0 : len(val)/2]), util.ParseUint64(val[len(val)/2:]), true
	}

	return stone * 2024, 0, false
}

func Part02(input []string) string {
	stones := util.Map(strings.Split(input[0], " "), util.ParseUint64)

	return strconv.Itoa(int(smartBlink(stones, 75)))
}

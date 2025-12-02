package day01

import (
	"log"
	"strconv"

	"github.com/kingkero/adventofcode/2025/go/util"
)

func Part01(input []string) string {
	// start
	var position int64 = 50
	foundZeros := 0

	for i := range input {
		// direction for the dial
		// L -> -1, R -> 1
		var sign int64 = 1
		if input[i][:1] == "L" {
			sign = -1
		}

		amount, err := strconv.ParseInt(input[i][1:], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		var newPosition int64 = position + (sign * amount)

		for newPosition < 0 {
			newPosition += 100
		}

		for newPosition > 99 {
			newPosition -= 100
		}

		if newPosition == 0 {
			foundZeros++
		}

		position = newPosition
	}

	return strconv.Itoa(foundZeros)
}

func Part02(input []string) string {
	position := 50
	zeroCount := 0

	for i := range input {
		distance := util.ParseInt(input[i][1:])

		step := 1
		if input[i][:1] == "L" {
			step = -1
		}

		for i := 0; i < distance; i++ {
			position = (position + step) % 100
			if position < 0 {
				position += 100
			}

			if position == 0 {
				zeroCount++
			}
		}
	}

	return strconv.Itoa(zeroCount)
}

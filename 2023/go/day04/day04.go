package day04

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"

	"github.com/kingkero/adventofcode/2023/go/util"
)

func part01(lines []string) int {
	result := 0

	whiteSpaceRegex := regexp.MustCompile("\\s+")

	for _, line := range lines {
		parts := whiteSpaceRegex.Split(line, -1)

		// cardNumber := parts[1]
		var winners []int
		amountOfWinners := 0

		isWinner := true
		for i := 2; i < len(parts); i++ {
			if parts[i] == "|" {
				isWinner = false
				continue
			}

			number, _ := strconv.ParseInt(parts[i], 10, 64)
			if isWinner {
				winners = append(winners, int(number))
				continue
			}

			if util.IntsContains(winners, int(number)) {
				amountOfWinners++
			}
		}

		if amountOfWinners > 0 {
			result += int(math.Pow(2.0, float64(amountOfWinners-1)))
		}
	}

	return result
}

func part02(lines []string) int {
	result := 0

	whiteSpaceRegex := regexp.MustCompile("\\s+")
	cardToWinners := make(map[int]int)

	for _, line := range lines {
		parts := whiteSpaceRegex.Split(line, -1)

		cardNumber, _ := strconv.ParseInt(parts[1][:len(parts[1])-1], 10, 64)
		cardNumberInt := int(cardNumber)
		var winners []int
		amountOfWinners := 0

		isWinner := true
		for i := 2; i < len(parts); i++ {
			if parts[i] == "|" {
				isWinner = false
				continue
			}

			number, _ := strconv.ParseInt(parts[i], 10, 64)
			if isWinner {
				winners = append(winners, int(number))
				continue
			}

			if util.IntsContains(winners, int(number)) {
				amountOfWinners++
			}
		}

		cardToWinners[cardNumberInt] = amountOfWinners
	}

	copies := make(map[int]int)
	for card := range cardToWinners {
		copies[card] = 1
	}

	for card, winners := range cardToWinners {
		for i := 0; i < copies[card]; i++ {
			for j := 0; j < winners; j++ {
				copies[card+j+1]++
			}
		}
	}

	fmt.Println(copies)

	for a, b := range copies {
		result += b
		fmt.Printf("%d = %d\n", a, b)
	}

	return result
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	return part01(lines), part02(lines)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getColorMention(reveal string, color string) float64 {
	match := regexp.MustCompile("([0-9]+) " + color).FindStringSubmatch(reveal)
	if len(match) == 0 {
		return 0
	}

	result, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func solve() (int, int) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var resultPart01 int64 = 0
	var resultPart02 float64 = 0.0

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ": ")
		foundRed, foundBlue, foundGreen := 0.0, 0.0, 0.0

		id, _ := strconv.ParseInt(regexp.MustCompile("([0-9]+)$").FindString(split[0]), 10, 16)
		for _, reveal := range strings.Split(split[1], "; ") {
			foundRed = math.Max(foundRed, getColorMention(reveal, "red"))
			foundBlue = math.Max(foundBlue, getColorMention(reveal, "blue"))
			foundGreen = math.Max(foundGreen, getColorMention(reveal, "green"))
		}

		if foundRed <= 12 && foundBlue <= 14 && foundGreen <= 13 {
			resultPart01 += id
		}
		resultPart02 += foundRed * foundBlue * foundGreen

	}
	return int(resultPart01), int(resultPart02)
}

func main() {
	solution01, solution02 := solve()
	fmt.Println("Part 1: ", solution01)
	fmt.Println("Part 2: ", solution02)
}

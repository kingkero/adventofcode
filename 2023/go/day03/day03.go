package day03

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/kingkero/adventofcode/2023/go/util"
)

// Check a single line against its previous and next line if it contains candidate numbers.
func checkLine(prevLine, currentLine, nextLine string) []int {
	var result []int

	lineLength := len(currentLine)
	regex := regexp.MustCompile("[0-9]+")
	symbols := regex.Split(currentLine, -1)
	numbers := regex.FindAllString(currentLine, -1)
	leftDistance := len(symbols[0])
	symbolsIndex := 1

	removeDotRegex := regexp.MustCompile("(\\.|[0-9])+")
	for _, number := range numbers {
		checkForSymbol := ""

		rightBoundary := leftDistance + len(number)
		if rightBoundary+1 < lineLength {
			rightBoundary++
			checkForSymbol += currentLine[(leftDistance + len(number)):rightBoundary]
		}

		leftBoundary := leftDistance
		if leftDistance > 0 {
			leftBoundary--
			checkForSymbol += currentLine[leftBoundary:leftDistance]
		}

		if len(prevLine) > 0 {
			checkForSymbol += prevLine[leftBoundary:rightBoundary]
		}
		if len(nextLine) > 0 {
			checkForSymbol += nextLine[leftBoundary:rightBoundary]
		}

		if len(removeDotRegex.ReplaceAllString(checkForSymbol, "")) > 0 {
			parsed, _ := strconv.ParseInt(number, 10, 64)
			result = append(result, int(parsed))
		}

		leftDistance += len(number) + len(symbols[symbolsIndex])
		symbolsIndex++
	}

	return result
}

func part01(lines []string) int {
	result := 0

	prevLine := ""
	for i, line := range lines {
		nextLine := ""
		if i+1 < len(lines) {
			nextLine = lines[i+1]
		}
		importantLineNumbers := checkLine(prevLine, line, nextLine)
		result += util.SumInts(importantLineNumbers)

		prevLine = line
	}

	return result
}

func part02(lines []string) int {
	// TODO:
	// instead of splitting again and again each line and using checkForSymbol in part01
	// parse the whole file once into a proper readable format
	// so that it can be checked for:
	// - given number at position (line + col), are there adjacent symbols?
	// - given gear at position (line + col), get adjacent numbers
	result := 0

	asteriskRegex := regexp.MustCompile("\\*")

	// prevLine := ""
	for lineNumber, line := range lines {
		/*
			nextLine := ""
			if i+1 < len(lines) {
				nextLine = lines[i+1]
			}
		*/

		parts := asteriskRegex.Split(line, -1)
		if len(parts) < 2 {
			continue
		}

		leftDistance := len(parts[0])
		partsIndex := 1

		for gear := 0; gear < len(parts)-1; gear++ {
			fmt.Printf("gear at line %d col %d\n", lineNumber+1, leftDistance+1)

			leftDistance += len(parts[partsIndex])
			partsIndex++
		}

		// prevLine = line
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

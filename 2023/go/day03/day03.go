package day03

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Read a complete file line by line into memory.
func readFile(input string) ([]string, error) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

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

func sum(values []int) int {
	result := 0
	for _, val := range values {
		result += val
	}
	return result
}

func Solve(input string) (int, int) {
	lines, err := readFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	result := 0

	prevLine := ""
	for i, line := range lines {
		nextLine := ""
		if i+1 < len(lines) {
			nextLine = lines[i+1]
		}
		importantLineNumbers := checkLine(prevLine, line, nextLine)
		result += sum(importantLineNumbers)

		prevLine = line
	}

	return result, 0
}

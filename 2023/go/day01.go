package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part01() int {
	file, err := os.Open("./inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := 0

	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		charLength := len(chars)
		var foundLeft, foundRight int8 = -1, -1
		for i := 0; i < charLength; i++ {
			if foundLeft == -1 {
				leftNum, err := strconv.ParseInt(chars[i], 10, 8)
				if err == nil {
					foundLeft = int8(leftNum)
				}
			}
			if foundRight == -1 {
				rightNum, err := strconv.ParseInt(chars[charLength-1-i], 10, 8)
				if err == nil {
					foundRight = int8(rightNum)
				}
			}

			if foundLeft != -1 && foundRight != -1 {
				break
			}
		}

		result += int(foundLeft*10) + int(foundRight)
	}

	return result
}

func checkForNumber(stringVal string, intVal int, charLength int, chars []string, numbers *[]int, i int) {
	stringValChars := strings.Split(stringVal, "")
	if (i + len(stringValChars)) > charLength {
		return
	}

	for j := 0; j < len(stringValChars); j++ {
		if chars[i+j] != stringValChars[j] {
			return
		}
	}

	*numbers = append(*numbers, intVal)
}

func part02() int {
	file, err := os.Open("./inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := 0

	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for scanner.Scan() {
		numbers := []int{}
		chars := strings.Split(scanner.Text(), "")
		charLength := len(chars)
		// var foundLeft, foundRight int8 = -1, -1
		for i := 0; i < charLength; i++ {
			num, err := strconv.ParseInt(chars[i], 10, 8)
			if err == nil {
				numbers = append(numbers, int(num))
			}

			for stringVal, intVal := range numberMap {
				checkForNumber(stringVal, intVal, charLength, chars, &numbers, i)
			}
		}

		if len(numbers) == 1 {
			result += numbers[0]*10 + numbers[0]
		} else {
			result += numbers[0]*10 + numbers[len(numbers)-1]
		}
	}

	return result
}

func main() {
	fmt.Println("Part 1: ", part01())
	fmt.Println("Part 2: ", part02())
}

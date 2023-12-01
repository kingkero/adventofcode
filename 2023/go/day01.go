package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
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
		var foundLeft, foundRight int = -1, -1
		for i := 0; i < charLength; i++ {
			if foundLeft == -1 {
				leftNum, err := strconv.ParseInt(chars[i], 10, 32)
				if err == nil {
					foundLeft = int(leftNum)
				}
			}
			if foundRight == -1 {
				rightNum, err := strconv.ParseInt(chars[charLength-1-i], 10, 32)
				if err == nil {
					foundRight = int(rightNum)
				}
			}
		}

		result += (foundLeft * 10) + foundRight
	}

	fmt.Println(result)
}

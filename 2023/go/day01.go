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

	fmt.Println(result)
}

package util

import (
	"bufio"
	"log"
	"os"
)

// Read a complete file line by line into memory.
func ReadFile(input string) ([]string, error) {
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

// Get the sum of all ints in an array.
func SumInts(values []int) int {
	result := 0
	for _, val := range values {
		result += val
	}
	return result
}

func IntsContains(haystack []int, needle int) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
}

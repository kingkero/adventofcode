package util

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

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

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func Filter[T any](ts []T, f func(T) bool) []T {
	var result []T
	for _, t := range ts {
		if f(t) {
			result = append(result, t)
		}
	}
	return result
}

func Sum[T Number](ts []T) T {
	var sum T
	for i := range ts {
		sum += ts[i]
	}
	return sum
}

func ParseInt(value string) int {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(val)
}

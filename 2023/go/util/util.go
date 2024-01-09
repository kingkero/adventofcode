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

type Point struct {
	Row int
	Col int
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

// Return a list of elements, where the new element is the result of
// running the passed function.
func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

// Filter a list, return only elements that return true for the compare method.
func Filter[T any](ts []T, f func(T) bool) []T {
	var result []T
	for _, t := range ts {
		if f(t) {
			result = append(result, t)
		}
	}
	return result
}

// Sum a slice/array of numbers.
func Sum[T Number](ts []T) T {
	var sum T
	for i := range ts {
		sum += ts[i]
	}
	return sum
}

// Parse an integer, panic if an error happened.
func ParseInt(value string) int {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(val)
}

// Calculate Hamming distance between two strings.
// see https://en.wikipedia.org/wiki/Hamming_distance
func Hamming(a, b string) int {
	if len(a) != len(b) {
		return -1
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance
}

func SumOccurrences(data [][]int) map[int]int {
	result := make(map[int]int)

	for _, list := range data {
		for _, element := range list {
			result[element]++
		}
	}

	return result
}

func GetMapKeys[K comparable, V any](data map[K]V) []K {
	result := make([]K, len(data))
	position := 0
	for key := range data {
		result[position] = key
		position++
	}
	return result
}

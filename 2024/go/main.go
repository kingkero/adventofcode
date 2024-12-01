package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

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

func part1(lines []string) {
	left := make([]uint, 0, len(lines))
	right := make([]uint, 0, len(lines))

	for _, line := range lines {
		var l, r uint
		_, err := fmt.Sscanf(line, "%d   %d", &l, &r)
		if err != nil {
			log.Fatal(line, err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	distances := uint(0)
	for i := 0; i < len(left); i++ {
		if right[i] < left[i] {
			distances += left[i] - right[i]
		} else {
			distances += right[i] - left[i]
		}
	}

	fmt.Printf("%v\n", distances)
}

func part2(lines []string) {
	left := make([]uint, 0, len(lines))
	right := make([]uint, 0, len(lines))
	similarities := make(map[uint]uint)

	for _, line := range lines {
		var l, r uint
		_, err := fmt.Sscanf(line, "%d   %d", &l, &r)
		if err != nil {
			log.Fatal(line, err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	result := uint(0)
	for i := 0; i < len(left); i++ {
		if v, ok := similarities[left[i]]; ok {
			result += left[i] * v
			continue
		}

		countSimilarities := uint(0)
		for _, item := range right {
			if item == left[i] {
				countSimilarities++
			}
		}
		similarities[left[i]] = countSimilarities
		result += left[i] * countSimilarities
	}

	fmt.Printf("%v\n", result)
}

func main() {
	data, _ := ReadFile("./inputs/day1_1.txt")

	part1(data)
	part2(data)
}

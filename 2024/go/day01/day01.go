package day01

import (
	"fmt"
	"log"
	"slices"
	"strconv"
)

func Part01(input []string) string {
	left := make([]uint, 0, len(input))
	right := make([]uint, 0, len(input))

	for _, line := range input {
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

	return strconv.Itoa(int(distances))
}

func Part02(input []string) string {
	left := make([]uint, 0, len(input))
	right := make([]uint, 0, len(input))
	similarities := make(map[uint]uint)

	for _, line := range input {
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

	return strconv.Itoa(int(result))
}

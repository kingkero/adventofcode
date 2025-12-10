package day09

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2025/go/util"
)

type point struct {
	x, y int
}

func (p point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p point) area(a point) int {
	firstPart := p.x - a.x
	if a.x > p.x {
		firstPart = a.x - p.x
	}

	secondPart := p.y - a.y
	if a.y > p.y {
		secondPart = a.y - p.y
	}

	return (firstPart + 1) * (secondPart + 1)
}

func Part01(input []string) string {
	points := make([]point, len(input))
	areas := make([]int, 0)

	for i, line := range input {
		coordinates := util.Map(strings.Split(line, ","), util.ParseInt)
		points[i] = point{coordinates[0], coordinates[1]}

		for j := 0; j < i; j++ {
			areas = append(areas, points[i].area(points[j]))
		}
	}

	slices.Sort(areas)
	slices.Reverse(areas)

	return strconv.Itoa(areas[0])
}

func Part02(input []string) string {
	return strconv.Itoa(0)
}

package day08

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2025/go/util"
)

type point struct {
	x, y, z int
}

func (a point) String() string {
	return fmt.Sprintf("(%d, %d, %d)", a.x, a.y, a.z)
}

// just an approximation
// skips the sqrt, since sqrt(x) < sqrt(y) <=> x < y
func (a point) distanceTo(b point) int {
	xDistance := a.x - b.x
	yDistance := a.y - b.y
	zDistance := a.z - b.z

	return xDistance*xDistance + yDistance*yDistance + zDistance*zDistance
}

type circuit struct {
	pointIndices map[int]bool
}

func (c circuit) contains(pointIndex int) bool {
	_, ok := c.pointIndices[pointIndex]
	return ok
}

type cacheKey struct {
	a, b int
}

func Part01(input []string) string {
	points := make([]point, len(input))

	// a single circuit contains the index of its points
	// circuits := make([][]int, 0)

	// distanceCache := make(map[cacheKey]int)

	distances := make([]int, 0)
	distanceToPoints := make(map[int][2]int)

	for i, line := range input {
		coordinates := util.Map(strings.Split(line, ","), util.ParseInt)
		points[i] = point{coordinates[0], coordinates[1], coordinates[2]}

		for j := 0; j < i; j++ {
			distance := points[j].distanceTo(points[i])

			distances = append(distances, distance)
			distanceToPoints[distance] = [2]int{j, i}
			// distanceCache[cacheKey{j, i}] = points[j].distanceTo(points[i])
		}
	}

	slices.Sort(distances)

	circuits := make([]circuit, 0, len(points))

OUTER:
	for _, distance := range distances[:1_000] {
		parts := distanceToPoints[distance]

		inCircuitIds := make([]int, 0)

		for i, circuit := range circuits {
			if circuit.contains(parts[0]) && circuit.contains(parts[1]) {
				// already got it
				continue OUTER
			}

			if circuit.contains(parts[0]) {
				inCircuitIds = append(inCircuitIds, i)
			} else if circuit.contains(parts[1]) {
				inCircuitIds = append(inCircuitIds, i)
			}
		}

		if len(inCircuitIds) == 0 {
			circuits = append(circuits, circuit{
				pointIndices: map[int]bool{
					parts[0]: true,
					parts[1]: true,
				},
			})

			continue OUTER
		}

		circuits[inCircuitIds[0]].pointIndices[parts[0]] = true
		circuits[inCircuitIds[0]].pointIndices[parts[1]] = true

		if len(inCircuitIds) == 1 {
			continue OUTER
		}

		for pointIndex := range circuits[inCircuitIds[1]].pointIndices {
			circuits[inCircuitIds[0]].pointIndices[pointIndex] = true
		}
		circuits = slices.Delete(circuits, inCircuitIds[1], inCircuitIds[1]+1)
	}

	sizes := make([]int, len(circuits))
	for _, circuit := range circuits {
		sizes = append(sizes, len(circuit.pointIndices))
	}

	slices.Sort(sizes)
	slices.Reverse(sizes)

	return strconv.Itoa(sizes[0] * sizes[1] * sizes[2])
}

func Part02(input []string) string {
	return strconv.Itoa(0)
}

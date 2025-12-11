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

type edge struct {
	p1, p2 point
}

func (e edge) isVertical() bool {
	return e.p1.x == e.p2.x
}

func (e edge) isPointOnEdge(p point) bool {
	if e.isVertical() {
		x := e.p1.x
		y1, y2 := min(e.p1.y, e.p2.y), max(e.p1.y, e.p2.y)
		return p.x == x && p.y >= y1 && p.y <= y2
	}

	y := e.p1.y
	x1, x2 := min(e.p1.x, e.p2.x), max(e.p1.x, e.p2.x)
	return p.y == y && p.x >= x1 && p.x <= x2
}

func isInsidePolygon(p point, edges []edge) bool {
	crossings := 0
	for _, e := range edges {
		if e.isVertical() {
			x := e.p1.x
			y1, y2 := min(e.p1.y, e.p2.y), max(e.p1.y, e.p2.y)
			if x > p.x && y1 <= p.y && p.y < y2 {
				crossings++
			}
		}
	}
	return crossings%2 == 1
}

func isInsideOrOnBoundary(p point, edges []edge) bool {
	// Check if on boundary first
	for _, e := range edges {
		if e.isPointOnEdge(p) {
			return true
		}
	}
	return isInsidePolygon(p, edges)
}

func isRectangleValid(p1, p2 point, edges []edge) bool {
	x1, x2 := min(p1.x, p2.x), max(p1.x, p2.x)
	y1, y2 := min(p1.y, p2.y), max(p1.y, p2.y)

	xSet := make(map[int]bool)
	ySet := make(map[int]bool)
	xSet[x1] = true
	xSet[x2] = true
	ySet[y1] = true
	ySet[y2] = true

	for _, e := range edges {
		if e.isVertical() {
			x := e.p1.x
			if x1 <= x && x <= x2 {
				xSet[x] = true
			}
			ey1, ey2 := min(e.p1.y, e.p2.y), max(e.p1.y, e.p2.y)
			if y1 <= ey1 && ey1 <= y2 {
				ySet[ey1] = true
			}
			if y1 <= ey2 && ey2 <= y2 {
				ySet[ey2] = true
			}
		} else {
			y := e.p1.y
			if y1 <= y && y <= y2 {
				ySet[y] = true
			}
			ex1, ex2 := min(e.p1.x, e.p2.x), max(e.p1.x, e.p2.x)
			if x1 <= ex1 && ex1 <= x2 {
				xSet[ex1] = true
			}
			if x1 <= ex2 && ex2 <= x2 {
				xSet[ex2] = true
			}
		}
	}

	xs := make([]int, 0, len(xSet))
	for x := range xSet {
		xs = append(xs, x)
	}
	slices.Sort(xs)

	ys := make([]int, 0, len(ySet))
	for y := range ySet {
		ys = append(ys, y)
	}
	slices.Sort(ys)

	for i := 0; i < len(xs)-1; i++ {
		for j := 0; j < len(ys)-1; j++ {
			cellWidth := xs[i+1] - xs[i]
			cellHeight := ys[j+1] - ys[j]

			// If cell has interior points, check one
			if cellWidth > 1 && cellHeight > 1 {
				tx := xs[i] + 1
				ty := ys[j] + 1
				if !isInsidePolygon(point{tx, ty}, edges) {
					return false
				}
			}
		}
	}

	for _, x := range xs {
		for _, y := range ys {
			if !isInsideOrOnBoundary(point{x, y}, edges) {
				return false
			}
		}
	}

	return true
}

func Part02(input []string) string {
	points := make([]point, len(input))
	for i, line := range input {
		coords := util.Map(strings.Split(line, ","), util.ParseInt)
		points[i] = point{coords[0], coords[1]}
	}

	edges := make([]edge, len(points))
	for i := 0; i < len(points); i++ {
		next := (i + 1) % len(points)
		edges[i] = edge{points[i], points[next]}
	}

	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if isRectangleValid(points[i], points[j], edges) {
				area := points[i].area(points[j])
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return strconv.Itoa(maxArea)
}

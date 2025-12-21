package day12

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Point struct{ r, c int }

type Orientation struct {
	points     []Point
	maxR, maxC int
}

type Region struct {
	width, height int
	counts        []int
}

func parseInput(input []string) ([][]*Orientation, []Region) {
	rawShapes := [][]string{}
	var regions []Region

	i := 0
	for i < len(input) {
		line := input[i]
		if strings.Contains(line, "x") {
			break
		}
		if strings.HasSuffix(line, ":") {
			i++
			shapeLines := []string{}
			for i < len(input) && input[i] != "" {
				if strings.HasSuffix(input[i], ":") || strings.Contains(input[i], "x") {
					break
				}
				shapeLines = append(shapeLines, input[i])
				i++
			}
			rawShapes = append(rawShapes, shapeLines)
			if i < len(input) && input[i] == "" {
				i++
			}
		} else {
			i++
		}
	}

	shapeOrientations := make([][]*Orientation, len(rawShapes))
	for idx, raw := range rawShapes {
		shapeOrientations[idx] = getAllOrientations(raw)
	}

	re := regexp.MustCompile(`(\d+)x(\d+):\s*(.*)`)
	for i < len(input) {
		line := input[i]
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			width, _ := strconv.Atoi(matches[1])
			height, _ := strconv.Atoi(matches[2])
			countsStr := strings.Fields(matches[3])
			counts := make([]int, len(countsStr))
			for j, s := range countsStr {
				counts[j], _ = strconv.Atoi(s)
			}
			regions = append(regions, Region{width, height, counts})
		}
		i++
	}

	return shapeOrientations, regions
}

func getAllOrientations(shapeLines []string) []*Orientation {
	rows := len(shapeLines)
	cols := 0
	for _, line := range shapeLines {
		if len(line) > cols {
			cols = len(line)
		}
	}

	grid := make([][]bool, rows)
	for r, line := range shapeLines {
		grid[r] = make([]bool, cols)
		for c, ch := range line {
			grid[r][c] = ch == '#'
		}
	}

	orientations := []*Orientation{}
	seen := make(map[string]bool)

	current := grid
	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			points := gridToPoints(current)
			key := pointsKey(points)
			if !seen[key] {
				seen[key] = true
				maxR, maxC := 0, 0
				for _, p := range points {
					if p.r > maxR {
						maxR = p.r
					}
					if p.c > maxC {
						maxC = p.c
					}
				}
				orientations = append(orientations, &Orientation{points, maxR, maxC})
			}
			current = rotate(current)
		}
		if flip == 0 {
			current = flipH(grid)
		}
	}

	return orientations
}

func gridToPoints(grid [][]bool) []Point {
	var points []Point
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] {
				points = append(points, Point{r, c})
			}
		}
	}
	if len(points) == 0 {
		return points
	}
	minR, minC := points[0].r, points[0].c
	for _, p := range points {
		if p.r < minR {
			minR = p.r
		}
		if p.c < minC {
			minC = p.c
		}
	}
	for i := range points {
		points[i].r -= minR
		points[i].c -= minC
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].r != points[j].r {
			return points[i].r < points[j].r
		}
		return points[i].c < points[j].c
	})
	return points
}

func pointsKey(points []Point) string {
	if len(points) == 0 {
		return ""
	}
	maxR, maxC := 0, 0
	for _, p := range points {
		if p.r > maxR {
			maxR = p.r
		}
		if p.c > maxC {
			maxC = p.c
		}
	}
	grid := make([][]byte, maxR+1)
	for i := range grid {
		grid[i] = make([]byte, maxC+1)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for _, p := range points {
		grid[p.r][p.c] = '#'
	}
	var sb strings.Builder
	for _, row := range grid {
		sb.Write(row)
		sb.WriteByte('\n')
	}
	return sb.String()
}

func rotate(grid [][]bool) [][]bool {
	if len(grid) == 0 {
		return grid
	}
	rows := len(grid)
	cols := len(grid[0])
	newGrid := make([][]bool, cols)
	for i := range newGrid {
		newGrid[i] = make([]bool, rows)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			newGrid[c][rows-1-r] = grid[r][c]
		}
	}
	return newGrid
}

func flipH(grid [][]bool) [][]bool {
	if len(grid) == 0 {
		return grid
	}
	rows := len(grid)
	cols := len(grid[0])
	newGrid := make([][]bool, rows)
	for r := 0; r < rows; r++ {
		newGrid[r] = make([]bool, cols)
		for c := 0; c < cols; c++ {
			newGrid[r][c] = grid[r][cols-1-c]
		}
	}
	return newGrid
}

// Bitmask operations for fast grid manipulation
func canPlaceBitmask(grid []uint64, width int, orient *Orientation, startR, startC int) bool {
	height := len(grid)
	// Quick bounds check
	if startR < 0 || startC < 0 || startR+orient.maxR >= height || startC+orient.maxC >= width {
		return false
	}
	for _, p := range orient.points {
		r, c := startR+p.r, startC+p.c
		if grid[r]&(1<<c) != 0 {
			return false
		}
	}
	return true
}

func placeBitmask(grid []uint64, points []Point, startR, startC int) {
	for _, p := range points {
		grid[startR+p.r] |= 1 << (startC + p.c)
	}
}

func unplaceBitmask(grid []uint64, points []Point, startR, startC int) {
	for _, p := range points {
		grid[startR+p.r] &^= 1 << (startC + p.c)
	}
}

// presentInfo tracks shape orientations and last placement position to avoid duplicate searches
type presentInfo struct {
	orientations []*Orientation
	size         int
	shapeID      int // to identify identical shapes
}

func solve(grid []uint64, width, height int, presents []presentInfo, idx int, lastPos int) bool {
	if idx >= len(presents) {
		return true
	}

	curr := presents[idx]
	// For identical consecutive shapes, start from lastPos to avoid permutation duplicates
	startPos := 0
	if idx > 0 && presents[idx-1].shapeID == curr.shapeID {
		startPos = lastPos
	}

	totalPositions := height * width
	for pos := startPos; pos < totalPositions; pos++ {
		r := pos / width
		c := pos % width
		for _, orient := range curr.orientations {
			if canPlaceBitmask(grid, width, orient, r, c) {
				placeBitmask(grid, orient.points, r, c)
				if solve(grid, width, height, presents, idx+1, pos) {
					return true
				}
				unplaceBitmask(grid, orient.points, r, c)
			}
		}
	}

	return false
}

func canFitRegion(shapeOrientations [][]*Orientation, region Region) bool {
	// Early check: total area of shapes vs grid area
	totalArea := 0
	for shapeIdx, count := range region.counts {
		if count > 0 && len(shapeOrientations[shapeIdx]) > 0 {
			totalArea += count * len(shapeOrientations[shapeIdx][0].points)
		}
	}
	gridArea := region.width * region.height
	if totalArea > gridArea {
		return false
	}

	// Build list of presents to place
	var presents []presentInfo

	for shapeIdx, count := range region.counts {
		if count > 0 && len(shapeOrientations[shapeIdx]) > 0 {
			size := len(shapeOrientations[shapeIdx][0].points)
			for i := 0; i < count; i++ {
				presents = append(presents, presentInfo{
					orientations: shapeOrientations[shapeIdx],
					size:         size,
					shapeID:      shapeIdx,
				})
			}
		}
	}

	if len(presents) == 0 {
		return true
	}

	// Sort by size descending (place larger pieces first), then by shapeID for grouping
	sort.Slice(presents, func(i, j int) bool {
		if presents[i].size != presents[j].size {
			return presents[i].size > presents[j].size
		}
		return presents[i].shapeID < presents[j].shapeID
	})

	// Create bitmask grid
	grid := make([]uint64, region.height)

	return solve(grid, region.width, region.height, presents, 0, 0)
}

func Part01(input []string) string {
	shapeOrientations, regions := parseInput(input)

	count := 0
	for _, region := range regions {
		if canFitRegion(shapeOrientations, region) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func Part02(input []string) string {
	return ""
}

package day08

import (
	"fmt"

	"github.com/kingkero/adventofcode/2024/go/util"
)

const Dot = int32('.')

func Part01(input []string) string {
	rows := len(input)
	cols := len(input[0])

	data := make(map[int32][]*util.Point)

	for row, line := range input {
		for col, char := range line {
			if char == Dot {
				continue
			}

			p := &util.Point{
				X: col,
				Y: row,
			}
			data[char] = append(data[char], p)
		}
	}

	antinodes := make([]*util.Point, 0, rows*cols)

	for _, points := range data {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				distance := points[i].DistanceTo(points[j])

				// WIP: there are 2 possible antinodes, but need to dynamically have the X+ / X- and Y+distance/Y-distance
				antinode1 := &util.Point{
					X: points[i].X + (2 * distance),
					Y: points[i].Y - (2 * distance),
				}
				antinode2 := &util.Point{
					X: points[j].X + (2 * distance),
					Y: points[j].Y - (2 * distance),
				}

				if antinode1.X >= 0 && antinode1.X < cols && antinode1.Y >= 0 && antinode1.Y < rows {
					antinodes = append(antinodes, antinode1)
				}
				if antinode2.X >= 0 && antinode2.X < cols && antinode2.Y >= 0 && antinode2.Y < rows {
					antinodes = append(antinodes, antinode2)
				}
			}
		}
	}

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", antinodes)
	return ""
}

func Part02(_ []string) string {
	return ""
}

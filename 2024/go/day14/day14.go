package day14

import (
	"fmt"
	"strconv"

	"github.com/kingkero/adventofcode/2024/go/util"
)

type Robot struct {
	position util.Point
	velocity util.Point
}

func (r *Robot) move(maxWidth, maxHeight int) {
	r.position.X += r.velocity.X
	r.position.Y += r.velocity.Y

	if r.position.X < 0 {
		r.position.X += maxWidth
	}
	if r.position.X >= maxWidth {
		r.position.X -= maxWidth
	}

	if r.position.Y < 0 {
		r.position.Y += maxHeight
	}
	if r.position.Y >= maxHeight {
		r.position.Y -= maxHeight
	}
}

func Part01(input []string) string {
	width := 101
	middleWidth := (width - 1) / 2
	height := 103
	middleHeight := (height - 1) / 2

	var posX, posY, velX, velY int
	robotsInQuadrants := make([]int, 4)

	for _, line := range input {
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &posX, &posY, &velX, &velY)

		robot := &Robot{
			position: util.Point{X: posX, Y: posY},
			velocity: util.Point{X: velX, Y: velY},
		}

		for i := 0; i < 100; i++ {
			robot.move(width, height)
		}

		// skip robots on middle line
		if robot.position.X == middleWidth || robot.position.Y == middleHeight {
			continue
		}

		if robot.position.X < middleWidth && robot.position.Y < middleHeight {
			robotsInQuadrants[0]++
		}
		if robot.position.X < middleWidth && robot.position.Y > middleHeight {
			robotsInQuadrants[1]++
		}
		if robot.position.X > middleWidth && robot.position.Y > middleHeight {
			robotsInQuadrants[2]++
		}
		if robot.position.X > middleWidth && robot.position.Y < middleHeight {
			robotsInQuadrants[3]++
		}
	}

	return strconv.Itoa(robotsInQuadrants[0] * robotsInQuadrants[1] * robotsInQuadrants[2] * robotsInQuadrants[3])
}

func Part02(_ []string) string {
	return ""
}

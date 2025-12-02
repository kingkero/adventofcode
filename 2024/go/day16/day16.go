package day16

import (
	"fmt"
	"sync"

	"github.com/kingkero/adventofcode/2024/go/util"
)

type Field uint8

const (
	Fence Field = '#'
	Start Field = 'S'
	End   Field = 'E'
	Free  Field = '.'
)

type Matrix struct {
	rows  int
	cols  int
	data  [][]Field
	start util.Point
	end   util.Point
}

type Solution struct {
	matrix           *Matrix
	points           int
	visited          [][]bool
	current          util.Point
	currentDirection util.Direction
}

func Part01(input []string) string {
	matrix := &Matrix{
		rows:  len(input),
		cols:  len(input[0]),
		start: util.Point{},
		end:   util.Point{},
	}
	matrix.data = make([][]Field, matrix.rows)

	for row, line := range input {
		matrix.data[row] = make([]Field, matrix.cols)

		for col, char := range line {
			value := Field(char)
			matrix.data[row][col] = value

			if value == Start {
				matrix.start = util.Point{X: row, Y: col}
			} else if value == End {
				matrix.end = util.Point{X: row, Y: col}
			}
		}
	}

	solution1 := &Solution{
		matrix:           matrix,
		points:           0,
		visited:          make([][]bool, matrix.rows),
		current:          util.Point{X: matrix.start.X, Y: matrix.start.Y},
		currentDirection: util.DirectionEast,
	}
	for i := 0; i < matrix.rows; i++ {
		solution1.visited[i] = make([]bool, matrix.cols)
	}

	// move same direction until wall, spawn child processes for openings

	var solutions []int
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		c := make(chan int)
		go solution1.move(c)

		x := <-c
		solutions = append(solutions, x)
	}()
	waitGroup.Wait()
	fmt.Printf("%v\n", solutions)

	/*
		fmt.Printf("Matrix: %v\n", matrix)
		fmt.Printf("Solution1: %v\n", solution1)

		fmt.Printf("Solutions: %v\n", solutions)
	*/
	return ""
}

func (s *Solution) move(c chan int) {
	s.visited[s.current.X][s.current.Y] = true

	deltaX := 0
	deltaY := 0
	switch s.currentDirection {
	case util.DirectionNorth:
		deltaX = -1
		deltaY = 0
	case util.DirectionEast:
		deltaX = 0
		deltaY = 1
	case util.DirectionSouth:
		deltaX = 1
		deltaY = 0
	case util.DirectionWest:
		deltaX = 0
		deltaY = -1
	default:
		panic("unhandled default case")
	}

	if s.visited[s.current.X+deltaX][s.current.Y+deltaY] {
		close(c)
		return
	}

	// for direction east:
	// maybe spawn top
	// maybe spawn bottom

	if s.matrix.data[s.current.X+deltaX][s.current.Y+deltaY] == Fence {
		close(c)
		return
	}

	s.points++
	s.current.X += deltaX
	s.current.Y += deltaY

	if s.matrix.data[s.current.X+deltaX][s.current.Y+deltaY] == End {
		c <- s.points
		return
	}

	s.move(c)
}

func Part02(_ []string) string {
	return ""
}

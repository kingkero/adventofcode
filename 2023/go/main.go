package main

import (
	"fmt"
	"time"

	"github.com/kingkero/adventofcode/2023/go/day01"
	"github.com/kingkero/adventofcode/2023/go/day02"
	"github.com/kingkero/adventofcode/2023/go/day03"
	"github.com/kingkero/adventofcode/2023/go/day04"
	"github.com/kingkero/adventofcode/2023/go/day05"
	"github.com/kingkero/adventofcode/2023/go/day06"
)

func main() {
	start := time.Now()
	p01, p02 := day01.Solve("./day01/input.txt")
	fmt.Printf("Day 1:\t%v\t/\t%v\ttook %v\n", p01, p02, time.Since(start))

	start = time.Now()
	p01, p02 = day02.Solve("./day02/input.txt")
	fmt.Printf("Day 2:\t%v\t/\t%v\ttook %v\n", p01, p02, time.Since(start))

	start = time.Now()
	p01, p02 = day03.Solve("./day03/input.txt")
	fmt.Printf("Day 3:\t%v\t/\t%v\ttook %v\n", p01, p02, time.Since(start))

	start = time.Now()
	p01, p02 = day04.Solve("./day04/input.txt")
	fmt.Printf("Day 4:\t%v\t/\t%v\ttook %v\n", p01, p02, time.Since(start))

	start = time.Now()
	p01, p02 = day05.Solve("./day05/input.txt")
	fmt.Printf("Day 5:\t%v\t/\t%v\ttook %v\n", p01, p02, time.Since(start))

	start = time.Now()
	p01, p02 = day06.Solve("./day06/input.txt")
	fmt.Printf("Day 6:\t%v\t/\t%v\ttook %v\n", p01, p02, time.Since(start))
}

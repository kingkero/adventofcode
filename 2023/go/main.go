package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/kingkero/adventofcode/2023/go/day01"
	"github.com/kingkero/adventofcode/2023/go/day02"
	"github.com/kingkero/adventofcode/2023/go/day03"
	"github.com/kingkero/adventofcode/2023/go/day04"
	"github.com/kingkero/adventofcode/2023/go/day05"
	"github.com/kingkero/adventofcode/2023/go/day06"
	"github.com/kingkero/adventofcode/2023/go/day07"
	"github.com/kingkero/adventofcode/2023/go/day08"
	"github.com/kingkero/adventofcode/2023/go/day09"
	"github.com/kingkero/adventofcode/2023/go/day10"
	"github.com/kingkero/adventofcode/2023/go/day11"
	"github.com/kingkero/adventofcode/2023/go/day12"
	"github.com/kingkero/adventofcode/2023/go/day13"
	"github.com/kingkero/adventofcode/2023/go/day14"
	"github.com/kingkero/adventofcode/2023/go/day15"
	"github.com/kingkero/adventofcode/2023/go/day16"
	"github.com/kingkero/adventofcode/2023/go/day17"
	"github.com/kingkero/adventofcode/2023/go/day18"
	"github.com/kingkero/adventofcode/2023/go/day19"
	"github.com/kingkero/adventofcode/2023/go/day20"
	"github.com/kingkero/adventofcode/2023/go/day21"
	"github.com/kingkero/adventofcode/2023/go/day22"
)

func getValueWithLeadingZeroes(value int) string {
	prefix := ""
	if value < 10 {
		prefix = "0"
	}

	return prefix + strconv.Itoa(value)
}

type Solver func(file string) (int, int)

func writeSolvers(writer *tabwriter.Writer, solvers ...Solver) {
	previous := time.Now()
	p01, p02 := 0, 0

	for i := len(solvers) - 1; i >= 0; i-- {
		name := "day" + getValueWithLeadingZeroes(i+1)

		p01, p02 = solvers[i]("./" + name + "/input.txt")
		fmt.Fprintf(writer, "Day %d:\t%v\t/\t%v\ttook %v\n", i+1, p01, p02, time.Since(previous))
		previous = time.Now()
	}
}

func main() {
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	writeSolvers(
		writer,
		day01.Solve,
		day02.Solve,
		day03.Solve,
		day04.Solve,
		day05.Solve,
		day06.Solve,
		day07.Solve,
		day08.Solve,
		day09.Solve,
		day10.Solve,
		day11.Solve,
		day12.Solve,
		day13.Solve,
		day14.Solve,
		day15.Solve,
		day16.Solve,
		day17.Solve,
		day18.Solve,
		day19.Solve,
		day20.Solve,
		day21.Solve,
		day22.Solve,
	)

	writer.Flush()
}

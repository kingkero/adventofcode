package main

import (
	"fmt"
	"os"
	"slices"
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
	"github.com/kingkero/adventofcode/2023/go/day23"
	"github.com/kingkero/adventofcode/2023/go/day24"
	"github.com/kingkero/adventofcode/2023/go/day25"
	"github.com/kingkero/adventofcode/2023/go/util"
)

func getValueWithLeadingZeroes(value int) string {
	prefix := ""
	if value < 10 {
		prefix = "0"
	}

	return prefix + strconv.Itoa(value)
}

type Solver func(file string) (int, int)

func writeSolvers(writer *tabwriter.Writer, solvers map[int]Solver) {
	days := util.GetMapKeys(solvers)
	slices.Sort(days)
	slices.Reverse(days)

	previous := time.Now()
	p01, p02 := 0, 0
	for _, day := range days {
		name := "day" + getValueWithLeadingZeroes(day)

		p01, p02 = solvers[day]("./" + name + "/input.txt")
		fmt.Fprintf(writer, "Day %d:\t%v\t/\t%v\ttook %v\n", day, p01, p02, time.Since(previous))
		previous = time.Now()
	}
}

func main() {
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	solvers := map[int]Solver{
		1:  day01.Solve,
		2:  day02.Solve,
		3:  day03.Solve,
		4:  day04.Solve,
		5:  day05.Solve,
		6:  day06.Solve,
		7:  day07.Solve,
		8:  day08.Solve,
		9:  day09.Solve,
		10: day10.Solve,
		11: day11.Solve,
		12: day12.Solve,
		13: day13.Solve,
		14: day14.Solve,
		15: day15.Solve,
		16: day16.Solve,
		17: day17.Solve,
		18: day18.Solve,
		19: day19.Solve,
		20: day20.Solve,
		21: day21.Solve,
		22: day22.Solve,
		23: day23.Solve,
		24: day24.Solve,
		25: day25.Solve,
	}

	// day := flag.Int("day", -1, "which day to run")
	// flag.Parse()

	/*
		if *day > 0 {
			var newSolvers []Solver
			newSolvers = append(newSolvers, solvers[*day-1])
			solvers = newSolvers
		}
	*/

	writeSolvers(
		writer,
		solvers,
	)

	writer.Flush()
}

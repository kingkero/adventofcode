package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/echojc/aocutil"

	"github.com/kingkero/adventofcode/2024/go/day01"
	"github.com/kingkero/adventofcode/2024/go/day02"
	"github.com/kingkero/adventofcode/2024/go/day03"
	"github.com/kingkero/adventofcode/2024/go/day04"
	"github.com/kingkero/adventofcode/2024/go/day05"
	"github.com/kingkero/adventofcode/2024/go/day06"
	"github.com/kingkero/adventofcode/2024/go/day07"
	"github.com/kingkero/adventofcode/2024/go/day08"
	"github.com/kingkero/adventofcode/2024/go/day09"
	"github.com/kingkero/adventofcode/2024/go/day10"
	"github.com/kingkero/adventofcode/2024/go/day11"
	"github.com/kingkero/adventofcode/2024/go/day12"
	"github.com/kingkero/adventofcode/2024/go/day13"
	"github.com/kingkero/adventofcode/2024/go/day14"
	"github.com/kingkero/adventofcode/2024/go/day15"
	"github.com/kingkero/adventofcode/2024/go/day16"
)

type Day struct {
	Part01 func([]string) string
	Part02 func([]string) string
}

func main() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)

	for key, day := range []Day{
		{day01.Part01, day01.Part02},
		{day02.Part01, day02.Part02},
		{day03.Part01, day03.Part02},
		{day04.Part01, day04.Part02},
		{day05.Part01, day05.Part02},
		{day06.Part01, day06.Part02},
		{day07.Part01, day07.Part02},
		{day08.Part01, day08.Part02},
		{day09.Part01, day09.Part02},
		{day10.Part01, day10.Part02},
		{day11.Part01, day11.Part02},
		{day12.Part01, day12.Part02},
		{day13.Part01, day13.Part02},
		{day14.Part01, day14.Part02},
		{day15.Part01, day15.Part02},
		{day16.Part01, day16.Part02},
	} {
		solveDay(i, writer, key+1, day.Part01, day.Part02)
	}

	if err = writer.Flush(); err != nil {
		log.Fatal(err)
	}
}

func solveDay(i *aocutil.Input, writer *tabwriter.Writer, day int, part01 func([]string) string, part02 func([]string) string) {
	lines, err := i.Strings(2024, day)
	if err != nil {
		log.Fatal(err)
	}

	previous := time.Now()
	if _, err := fmt.Fprintf(writer, "Day %d:\t%v\t%v\ttook %v\n", day, part01(lines), part02(lines), time.Since(previous)); err != nil {
		log.Fatal(err)
	}
}

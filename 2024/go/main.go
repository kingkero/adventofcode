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
)

func main() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)

	solveDay(i, writer, 1, day01.Part01, day01.Part02)
	solveDay(i, writer, 2, day02.Part01, day02.Part02)
	solveDay(i, writer, 3, day03.Part01, day03.Part02)
	solveDay(i, writer, 4, day04.Part01, day04.Part02)

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

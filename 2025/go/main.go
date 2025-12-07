package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/echojc/aocutil"

	"github.com/kingkero/adventofcode/2025/go/day01"
	"github.com/kingkero/adventofcode/2025/go/day02"
	"github.com/kingkero/adventofcode/2025/go/day03"
	"github.com/kingkero/adventofcode/2025/go/day04"
	"github.com/kingkero/adventofcode/2025/go/day05"
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
	} {
		solveDay(i, writer, key+1, day.Part01, day.Part02)
	}

	if err = writer.Flush(); err != nil {
		log.Fatal(err)
	}
}

func solveDay(i *aocutil.Input, writer *tabwriter.Writer, day int, part01 func([]string) string, part02 func([]string) string) {
	lines, err := i.Strings(2025, day)
	if err != nil {
		log.Fatal(err)
	}

	previous := time.Now()
	if _, err := fmt.Fprintf(writer, "Day %d:\t%v\t%v\ttook %v\n", day, part01(lines), part02(lines), time.Since(previous)); err != nil {
		log.Fatal(err)
	}
}

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
)

func main() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	solveDay(i, writer, 1, day01.Part01, day01.Part02)
	solveDay(i, writer, 2, day02.Part01, day02.Part02)

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
	if _, err := fmt.Fprintf(writer, "Day %d / Part %d:\t%v\ttook %v\n", day, 1, part01(lines), time.Since(previous)); err != nil {
		log.Fatal(err)
	}

	previous = time.Now()
	if _, err := fmt.Fprintf(writer, "Day %d / Part %d:\t%v\ttook %v\n", day, 2, part02(lines), time.Since(previous)); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/echojc/aocutil"

	"github.com/kingkero/adventofcode/2024/go/day01"
)

/*
func part1(lines []string) {
	safeReports := 0

out:
	for _, line := range lines {
		// isAscending := true
		recordings := Map(strings.Split(line, " "), ParseInt)

		isAscending := recordings[0] < recordings[1]
		for i := 0; i < len(recordings)-1; i++ {
			if recordings[i] == recordings[i+1] {
				continue out
			}

			if isAscending {
				if recordings[i] > recordings[i+1] {
					continue out
				}

				if recordings[i+1]-recordings[i] > 3 {
					continue out
				}
			} else {
				if recordings[i] < recordings[i+1] {
					continue out
				}

				if recordings[i]-recordings[i+1] > 3 {
					continue out
				}
			}
		}

		safeReports++
	}

	fmt.Println(safeReports)
}

func part2(lines []string) {
	safeReports := 0

	for _, line := range lines {
		// isAscending := true
		recordings := Map(strings.Split(line, " "), ParseInt)

		isAscending := recordings[0] < recordings[1]
		errors := 0
		for i := 0; i < len(recordings)-1; i++ {
			if recordings[i] == recordings[i+1] {
				errors++
				continue
			}

			if isAscending {
				if recordings[i] > recordings[i+1] {
					errors++
					continue
				}

				if recordings[i+1]-recordings[i] > 3 {
					errors++
					continue
				}
			} else {
				if recordings[i] < recordings[i+1] {
					errors++
					continue
				}

				if recordings[i]-recordings[i+1] > 3 {
					errors++
					continue
				}
			}
		}

		if errors > 2 {
			continue
		}

		safeReports++
	}

	fmt.Println(safeReports)
}
*/

func main() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	solveDay01(i, writer)

	if err = writer.Flush(); err != nil {
		log.Fatal(err)
	}
}

func solveDay01(i *aocutil.Input, writer *tabwriter.Writer) {
	lines, err := i.Strings(2024, 1)
	if err != nil {
		log.Fatal(err)
	}

	previous := time.Now()
	if _, err := fmt.Fprintf(writer, "Day %d / Part %d:\t%v\ttook %v\n", 1, 1, day01.Part01(lines), time.Since(previous)); err != nil {
		log.Fatal(err)
	}

	previous = time.Now()
	if _, err := fmt.Fprintf(writer, "Day %d / Part %d:\t%v\ttook %v\n", 1, 2, day01.Part02(lines), time.Since(previous)); err != nil {
		log.Fatal(err)
	}
}

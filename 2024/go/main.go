package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"

	"github.com/kingkero/adventofcode/2024/go/day01"
)

type Solver interface {
	Part1([]string) string
	Part2([]string) string
}

func ReadFile(input string) ([]string, error) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

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

func main() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	solveDay01(i)
}

func solveDay01(i *aocutil.Input) {
	lines, err := i.Strings(2024, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(day01.Part01(lines))
	fmt.Println(day01.Part02(lines))
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func ParseInt(value string) int {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(val)
}

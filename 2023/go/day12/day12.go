package day12

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/kingkero/adventofcode/2023/go/util"
)

type Record struct {
	springs       []string
	damagedGroups []int
}

func getRecords(lines []string) []*Record {
	var result []*Record

	for _, line := range lines {
		parts := strings.Split(line, " ")
		result = append(result, &Record{
			strings.Split(strings.Trim(parts[0], "."), ""),
			util.Map(strings.Split(parts[1], ","), util.ParseInt),
		})
	}

	return result
}

func (record *Record) canFitGroups(col, group int) bool {
	// final call of recursion
	if group >= len(record.damagedGroups) {
		for i := 1; col+i < len(record.springs); i++ {
			if record.springs[col+i] == "#" {
				return false
			}
		}
		return true
	}

	expandCol := record.damagedGroups[group]

	if col+expandCol > len(record.springs) {
		return false
	}

	if col+expandCol == len(record.springs) {
		return (group + 1) >= len(record.damagedGroups)
	}

	// can only contain "#" or "?"
	for i := 1; i < expandCol; i++ {
		if record.springs[col+i] == "." {
			return false
		}
	}

	// needs to end in "." or "?"
	if record.springs[col+expandCol] == "#" {
		return false
	}

	return record.canFitGroups(col+expandCol+1, group+1)
}

func (record *Record) getTotalCombinations() int {
	combinations := 1

	return combinations
}

func getPossibleCombinations(records []*Record, writer *tabwriter.Writer) int {
	return util.Sum(util.Map(records, func(record *Record) int {
		tmp := record.getTotalCombinations()
		fmt.Fprintf(writer, "%v\t\t%d\n", strings.Join(record.springs, ""), tmp)
		return tmp
		// return record.getTotalCombinations()
	}))
}

func part01(records []*Record) int {
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	result := getPossibleCombinations(records, writer)
	writer.Flush()
	return result
}

func part02() int {
	return 0
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	records := getRecords(lines)

	return part01(records), part02()
}

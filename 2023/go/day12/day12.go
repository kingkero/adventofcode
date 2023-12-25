package day12

import (
	"log"
	"strings"

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
			strings.Split(parts[0], ""),
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

func getPossibleCombinations(records []*Record) int {
	result := 0
	for _, record := range records {
		if record.canFitGroups(0, 0) {
			result++
		}
	}
	return result
}

func part01(records []*Record) int {
	return getPossibleCombinations(records)
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

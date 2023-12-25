package day12

import (
	"log"
	"strings"

	"github.com/gookit/goutil/dump"
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

func part01(records []*Record) int {
	dump.P(records)
	return 0
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

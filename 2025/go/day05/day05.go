package day05

import (
	"cmp"
	"slices"
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2025/go/util"
)

// is inclusive
type freshnessRange struct {
	start, end int
}

func (f freshnessRange) inRange(value int) bool {
	return value >= f.start && value <= f.end
}

func getFreshnessRange(line string) freshnessRange {
	parts := strings.Split(line, "-")

	return freshnessRange{
		start: util.ParseInt(parts[0]),
		end:   util.ParseInt(parts[1]),
	}
}

func Part01(input []string) string {
	freshnessRanges := make([]freshnessRange, 0, len(input))

	doChecks := false
	sumFresh := 0

	for _, line := range input {
		if line == "" {
			doChecks = true
			continue
		}

		if !doChecks {
			freshnessRanges = append(freshnessRanges, getFreshnessRange(line))
		} else {
			value := util.ParseInt(line)

			for _, freshnessRange := range freshnessRanges {
				if freshnessRange.inRange(value) {
					sumFresh++
					break
				}
			}
		}
	}

	return strconv.Itoa(sumFresh)
}

func Part02(input []string) string {
	freshnessRanges := make([]freshnessRange, 0, len(input))

	for _, line := range input {
		if line == "" {
			break
		}

		freshnessRanges = append(freshnessRanges, getFreshnessRange(line))
	}

	slices.SortFunc(freshnessRanges, func(a, b freshnessRange) int {
		return cmp.Or(
			cmp.Compare(a.start, b.start),
			cmp.Compare(a.end, b.end),
		)
	})

	finalRanges := make([]freshnessRange, 0, len(freshnessRanges))

	finalRanges = append(finalRanges, freshnessRanges[0])
	freshnessRanges = freshnessRanges[1:]

OUTER:
	for _, current := range freshnessRanges {
		currentFinal := len(finalRanges)
		notIncluded := 0

		for i, final := range finalRanges {
			if final.inRange(current.start) {
				if final.inRange(current.end) {
					continue OUTER
				}

				finalRanges[i].end = current.end
			} else if final.inRange(current.end) {
				finalRanges[i].start = current.start
			} else {
				notIncluded++
			}
		}

		if notIncluded == currentFinal {
			finalRanges = append(finalRanges, current)
		}
	}

	totalIdSum := 0
	for _, current := range finalRanges {
		totalIdSum += current.end - current.start + 1
	}

	return strconv.Itoa(totalIdSum)
}

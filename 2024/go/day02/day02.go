package day02

import (
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2024/go/util"
)

func isLineSafe(line []int) bool {
	isAscending := line[0] < line[1]
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			return false
		}

		if isAscending {
			if line[i] > line[i+1] {
				return false
			}

			if line[i+1]-line[i] > 3 {
				return false
			}
		} else {
			if line[i] < line[i+1] {
				return false
			}

			if line[i]-line[i+1] > 3 {
				return false
			}
		}
	}

	return true
}

func Part01(input []string) string {
	safeReports := 0

	for _, line := range input {
		recordings := util.Map(strings.Split(line, " "), util.ParseInt)

		if isLineSafe(recordings) {
			safeReports++
		}
	}

	return strconv.Itoa(safeReports)
}

func Part02(input []string) string {
	safeReports := 0

out:
	for _, line := range input {
		recordings := util.Map(strings.Split(line, " "), util.ParseInt)

		if isLineSafe(recordings) {
			safeReports++
			continue
		}

		// remove one element at a time and check if the line is safe
		for i := 0; i < len(recordings); i++ {
			removed := make([]int, 0, len(recordings)-1)
			removed = append(removed, recordings[:i]...)
			removed = append(removed, recordings[i+1:]...)

			if isLineSafe(removed) {
				safeReports++
				continue out
			}
		}
	}

	return strconv.Itoa(safeReports)
}

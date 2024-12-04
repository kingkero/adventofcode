package day02

import (
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2024/go/util"
)

func Part01(input []string) string {
	safeReports := 0

out:
	for _, line := range input {
		// isAscending := true
		recordings := util.Map(strings.Split(line, " "), util.ParseInt)

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

	return strconv.Itoa(int(safeReports))
}

func Part02(input []string) string {
	return ""
}

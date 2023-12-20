package day11

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

func isRowEmpty(row []string) bool {
	for col := 0; col < len(row); col++ {
		if row[col] != "." {
			return false
		}
	}
	return true
}

func isColEmpty(fields [][]string, col int) bool {
	for row := 0; row < len(fields); row++ {
		if fields[row][col] != "." {
			return false
		}
	}
	return true
}

func getExpandedFields(lines []string) [][]string {
	var expandRows, expandCols []int

	fields := make([][]string, len(lines))
	for row := 0; row < len(lines); row++ {
		fields[row] = strings.Split(lines[row], "")

		if isRowEmpty(fields[row]) {
			expandRows = append(expandRows, row+len(expandRows))
		}
	}

	for col := 0; col < len(fields[0]); col++ {
		if isColEmpty(fields, col) {
			expandCols = append(expandCols, col+len(expandCols))
		}
	}

	newRows := len(lines) + len(expandRows)
	newCols := len(lines[0]) + len(expandCols)

	expandedFields := make([][]string, newRows)

	referenceRow := 0
	for row := 0; row < newRows; row++ {
		newRow := make([]string, newCols)
		referenceCol := 0
		for col := 0; col < newCols; col++ {
			newRow[col] = fields[referenceRow][referenceCol]
			referenceCol++

			if slices.Contains(expandCols, col) {
				col++
				newRow[col] = "."
			}
		}
		expandedFields[row] = newRow
		referenceRow++

		if slices.Contains(expandRows, row) {
			row++
			expandedFields[row] = strings.Split(strings.Repeat(".", newCols), "")
		}
	}

	return expandedFields
}

func part01(lines []string) int {
	result := 0

	field := getExpandedFields(lines)

	for _, row := range field {
		fmt.Println(strings.Join(row, ""))
	}

	return result
}

func part02(lines []string) int {
	result := 0

	return result
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	return part01(lines), part02(lines)
}

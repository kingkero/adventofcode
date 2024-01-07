package day15

import (
	"log"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

func hash(data string) int {
	result := 0

	for _, element := range []rune(data) {
		result += int(element)
		result *= 17
		result = result % 256
	}

	return result
}

func part01(lines []string) int {
	sequences := strings.Split(lines[0], ",")
	hashed := util.Map(sequences, hash)
	return util.Sum(hashed)
}

func part02(lines []string) int {
	return 0
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	return part01(lines), part02(lines)
}

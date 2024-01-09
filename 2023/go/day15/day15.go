package day15

import (
	"log"
	"slices"
	"strings"

	"github.com/gookit/goutil/dump"
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
	return util.Sum(util.Map(strings.Split(lines[0], ","), hash))
}

type Lens struct {
	label       string
	focalLength int
}

func part02(lines []string) int {
	parts := strings.Split(lines[0], ",")
	hashMap := make(map[int][]*Lens)

	for _, part := range parts {
		equals := strings.Split(part, "=")

		if len(equals) == 1 {
			label := equals[0][:len(equals[0])-1]
			hashMap[hash(label)] = slices.DeleteFunc(hashMap[hash(label)], func(lens *Lens) bool {
				return lens.label == label
			})
			continue
		}

		lens := &Lens{equals[0], util.ParseInt(equals[1])}
		hashMap[hash(equals[0])] = append(hashMap[hash(equals[0])], lens)
	}

	dump.P(hashMap)

	return 0
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	return part01(lines), part02(lines)
}

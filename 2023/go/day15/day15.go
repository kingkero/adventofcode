package day15

import (
	"log"
	"slices"
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
			labelHash := hash(label)
			hashMap[labelHash] = slices.DeleteFunc(hashMap[labelHash], func(lens *Lens) bool {
				return lens.label == label
			})
			continue
		}

		label := equals[0]
		labelHash := hash(label)

		lens := &Lens{label, util.ParseInt(equals[1])}

		if index := slices.IndexFunc(hashMap[labelHash], func(lens *Lens) bool {
			return lens.label == label
		}); index != -1 {
			hashMap[labelHash][index] = lens
			continue
		}

		hashMap[hash(equals[0])] = append(hashMap[hash(equals[0])], lens)
	}

	result := 0
	for box, list := range hashMap {
		for index, lens := range list {
			result += (box + 1) * (index + 1) * lens.focalLength
		}
	}

	return result
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	return part01(lines), part02(lines)
}

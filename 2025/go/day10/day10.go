package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2025/go/util"
)

type button struct {
	targets []int
}

type lights []bool

func (l lights) equals(other lights) bool {
	if len(l) != len(other) {
		return false
	}

	for i, light := range l {
		if light != other[i] {
			return false
		}
	}

	return true
}

func (l lights) String() string {
	builder := strings.Builder{}
	for _, light := range l {
		if light {
			builder.WriteString("#")
		} else {
			builder.WriteString(".")
		}
	}

	return builder.String()
}

func Part01(input []string) string {
	for _, line := range input {
		parts := strings.Split(line, " ")

		// lights := make([]int, len(parts[0])-2)
		goal := make(lights, len(parts[0])-2)
		buttons := make([]button, len(parts)-2)

		for pos, light := range parts[0][1 : len(parts[0])-1] {
			goal[pos] = light == '#'
		}

		for pos, definition := range parts[1 : len(parts)-1] {
			buttons[pos] = button{
				targets: util.Map(strings.Split(strings.Trim(definition, "()"), ","), util.ParseInt),
			}
		}

		fmt.Println(goal)
		fmt.Printf("%#v\n", buttons)
	}

	return strconv.Itoa(0)
}

func Part02(input []string) string {
	return strconv.Itoa(0)
}

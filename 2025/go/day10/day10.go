package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2025/go/util"
)

func Part01(input []string) string {
	for _, line := range input {
		parts := strings.Split(line, " ")

		goal := 0
		buttons := make([]int, len(parts)-2)

		for pos, light := range parts[0][1 : len(parts[0])-1] {
			if light == '#' {
				goal |= 1 << pos
			}
		}

		for buttonIndex, definition := range parts[1 : len(parts)-1] {
			for _, pos := range util.Map(strings.Split(strings.Trim(definition, "()"), ","), util.ParseInt) {
				buttons[buttonIndex] |= 1 << pos
			}
		}

		fmt.Printf("%b\n", goal)
		for _, button := range buttons {
			fmt.Printf("%b\n", button)
		}
		fmt.Println("")

		// lights := make([]int, len(parts[0])-2)
		/*
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
		*/
	}

	return strconv.Itoa(0)
}

func Part02(input []string) string {
	return strconv.Itoa(0)
}

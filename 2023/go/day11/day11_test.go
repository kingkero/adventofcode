package day11

import "testing"

func TestSolutionDay11(t *testing.T) {
	solution01, solution02 := 10231178, 622120986954
	part01, part02 := Solve("./input.txt")

	if part01 != solution01 {
		t.Fatalf("Part01 solution expected %d, was %d", solution01, part01)
	}

	if part02 != solution02 {
		t.Fatalf("Part02 solution expected %d, was %d", solution02, part02)
	}
}

package day06

import "testing"

func TestSolutionDay06(t *testing.T) {
	solution01, solution02 := 140220, 39570185
	part01, part02 := Solve("./input.txt")

	if part01 != solution01 {
		t.Fatalf("Part01 solution expected %d, was %d", solution01, part01)
	}

	if part02 != solution02 {
		t.Fatalf("Part02 solution expected %d, was %d", solution02, part02)
	}
}

package day01

import "testing"

func TestSolutionDay01(t *testing.T) {
	part01, part02 := Solve("./input.txt")

	if part01 != 55208 {
		t.Fatalf("Part01 solution expected %d, was %d", 55208, part01)
	}

	if part02 != 54578 {
		t.Fatalf("Part02 solution expected %d, was %d", 54578, part02)
	}
}

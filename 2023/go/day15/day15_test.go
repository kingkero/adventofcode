package day15

import "testing"

func TestSolutionDay15DemoData(t *testing.T) {
	expected := []int{-1, -1}

	part01, part02 := Solve("./input_demo.txt")
	actual := []int{part01, part02}

	for i, expectation := range expected {
		if expectation != actual[i] {
			t.Fatalf("demo data part %d expected %d, was %d", i+1, expectation, actual[i])
		}
	}
}

func TestSolutionDay15RealData(t *testing.T) {
	expected := []int{-1, -1}

	part01, part02 := Solve("./input.txt")
	actual := []int{part01, part02}

	for i, expectation := range expected {
		if expectation != actual[i] {
			t.Fatalf("real data part %d expected %d, was %d", i+1, expectation, actual[i])
		}
	}
}

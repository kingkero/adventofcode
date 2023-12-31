package day13

import "testing"

func TestSolutionDay13DemoData(t *testing.T) {
	expected := []int{405, 400}

	part01, part02 := Solve("./input_demo.txt")
	actual := []int{part01, part02}

	for i, expectation := range expected {
		if expectation != actual[i] {
			t.Fatalf("demo data part %d expected %d, was %d", i+1, expectation, actual[i])
		}
	}
}

func TestSolutionDay13RealData(t *testing.T) {
	expected := []int{34918, 33054}

	part01, part02 := Solve("./input.txt")
	actual := []int{part01, part02}

	for i, expectation := range expected {
		if expectation != actual[i] {
			t.Fatalf("real data part %d expected %d, was %d", i+1, expectation, actual[i])
		}
	}
}

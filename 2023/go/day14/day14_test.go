package day14

import "testing"

func TestSolutionDay14DemoData(t *testing.T) {
	expected := []int{136, 64}

	part01, part02 := Solve("./input_demo.txt")
	actual := []int{part01, part02}

	for i, expectation := range expected {
		if expectation != actual[i] {
			t.Fatalf("demo data part %d expected %d, was %d", i+1, expectation, actual[i])
		}
	}
}

func TestSolutionDay14RealData(t *testing.T) {
	expected := []int{109654, 94876}

	part01, part02 := Solve("./input.txt")
	actual := []int{part01, part02}

	for i, expectation := range expected {
		if expectation != actual[i] {
			t.Fatalf("real data part %d expected %d, was %d", i+1, expectation, actual[i])
		}
	}
}

package util

import (
	"testing"
)

func TestParseInt(t *testing.T) {
	cases := []string{
		"5",
		"13",
		"432",
	}
	expectedResults := []int{
		5,
		13,
		432,
	}

	for i, test := range cases {
		result := ParseInt(test)
		if result != expectedResults[i] {
			t.Fatalf("Exptected \"%v\" to be %d, got %d instead", test, expectedResults[i], result)
		}
	}
}

func TestHammingDistance(t *testing.T) {
	cases := [][]string{
		{"foo", "foo"},
		{"foo", "bar"},
		{"bar", "baz"},
	}
	expectedResults := []int{
		0,
		3,
		1,
	}

	for i, test := range cases {
		result := Hamming(test[0], test[1])
		if result != expectedResults[i] {
			t.Fatalf("Expected distance between %v and %v to be %d, got %d instead", test[0], test[1], expectedResults[i], result)
		}
	}
}

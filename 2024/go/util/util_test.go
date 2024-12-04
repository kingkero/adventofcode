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

package day10

import "testing"

var testInput = []string{}

func TestPart01(t *testing.T) {
	result := Part01(testInput)

	if result != "0" {
		t.Errorf("got %s, want 0", result)
	}
}

func BenchmarkPart01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part01(testInput)
	}
}

/*
func TestPart02(t *testing.T) {
	result := Part02(testInput)

	if result != "40" {
		t.Errorf("got %s, want 40", result)
	}
}

func BenchmarkPart02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part02(testInput)
	}
}
*/

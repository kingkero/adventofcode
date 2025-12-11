package day09

import "testing"

var testInput = []string{
	"7,1",
	"11,1",
	"11,7",
	"9,7",
	"9,5",
	"2,5",
	"2,3",
	"7,3",
}

func TestPart01(t *testing.T) {
	result := Part01(testInput)

	if result != "50" {
		t.Errorf("got %s, want 50", result)
	}
}

func BenchmarkPart01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part01(testInput)
	}
}

func TestPart02(t *testing.T) {
	result := Part02(testInput)

	if result != "24" {
		t.Errorf("got %s, want 24", result)
	}
}

func BenchmarkPart02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part02(testInput)
	}
}

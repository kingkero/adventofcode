package day10

import "testing"

var testInput = []string{
	"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
	"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
	"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
}

func TestPart01(t *testing.T) {
	result := Part01(testInput)

	if result != "7" {
		t.Errorf("got %s, want 7", result)
	}
}

func BenchmarkPart01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part01(testInput)
	}
}

func TestPart02(t *testing.T) {
	result := Part02(testInput)

	if result != "33" {
		t.Errorf("got %s, want 33", result)
	}
}

func BenchmarkPart02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part02(testInput)
	}
}

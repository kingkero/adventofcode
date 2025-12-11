package day11

import "testing"

var testInput = []string{
	"aaa: you hhh",
	"you: bbb ccc",
	"bbb: ddd eee",
	"ccc: ddd eee fff",
	"ddd: ggg",
	"eee: out",
	"fff: out",
	"ggg: out",
	"hhh: ccc fff iii",
	"iii: out",
}

func TestPart01(t *testing.T) {
	result := Part01(testInput)

	if result != "5" {
		t.Errorf("got %s, want 5", result)
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

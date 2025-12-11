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

var testInput2 = []string{
	"svr: aaa bbb",
	"aaa: fft",
	"fft: ccc",
	"bbb: tty",
	"tty: ccc",
	"ccc: ddd eee",
	"ddd: hub",
	"hub: fff",
	"eee: dac",
	"dac: fff",
	"fff: ggg hhh",
	"ggg: out",
	"hhh: out",
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

func TestPart02(t *testing.T) {
	result := Part02(testInput2)

	if result != "2" {
		t.Errorf("got %s, want 2", result)
	}
}

func BenchmarkPart02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part02(testInput2)
	}
}

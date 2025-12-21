package day12

import "testing"

var testInput = []string{
	"0:",
	"###",
	"##.",
	"##.",
	"",
	"1:",
	"###",
	"##.",
	".##",
	"",
	"2:",
	".##",
	"###",
	"##.",
	"",
	"3:",
	"##.",
	"###",
	"##.",
	"",
	"4:",
	"###",
	"#..",
	"###",
	"",
	"5:",
	"###",
	".#.",
	"###",
	"",
	"4x4: 0 0 0 0 2 0",
	"12x5: 1 0 1 0 2 2",
	"12x5: 1 0 1 0 3 2",
}

func TestPart01(t *testing.T) {
	result := Part01(testInput)

	if result != "2" {
		t.Errorf("got %s, want 2", result)
	}
}

func BenchmarkPart01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part01(testInput)
	}
}

/*
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
*/

package day04

import "testing"

func TestPart01(t *testing.T) {
	result := Part01([]string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	})

	if result != "13" {
		t.Errorf("got %s, want 13", result)
	}
}

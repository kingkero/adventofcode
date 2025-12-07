package day06

import "testing"

func TestPart01(t *testing.T) {
	result := Part01([]string{
		"123 328  51 64",
		"45 64  387 23",
		"6 98  215 314",
		"*   +   *   +",
	})

	if result != "4277556" {
		t.Errorf("got %s, want 4277556", result)
	}
}

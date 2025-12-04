package day03

import "testing"

func TestPart01(t *testing.T) {
	result := Part01([]string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	})

	if result != "357" {
		t.Errorf("got %s, want 357", result)
	}
}

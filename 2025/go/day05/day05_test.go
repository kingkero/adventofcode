package day05

import "testing"

func TestPart01(t *testing.T) {
	result := Part01([]string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		"",
		"1",
		"5",
		"8",
		"11",
		"17",
		"32",
	})

	if result != "3" {
		t.Errorf("got %s, want 3", result)
	}
}

func TestPart02(t *testing.T) {
	result := Part02([]string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		"",
		"1",
		"5",
		"8",
		"11",
		"17",
		"32",
	})

	if result != "14" {
		t.Errorf("got %s, want 14", result)
	}
}

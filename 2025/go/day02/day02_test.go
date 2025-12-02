package day02

import "testing"

func TestIsInvalidID(t *testing.T) {
	examples := map[string]bool{
		"101":        false,
		"11":         true,
		"12":         false,
		"1111":       true,
		"1212":       true,
		"1112":       false,
		"1188511885": true,
		"1188511886": false,
	}

	for id, result := range examples {
		if isInvalidID(id) != result {
			t.Errorf("%s: expected %v, got %v", id, result, isInvalidID(id))
		}
	}
}

func TestPart01(t *testing.T) {
	result := Part01([]string{
		"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
	})

	if result != "1227775554" {
		t.Errorf("got %s, want 1227775554", result)
	}
}

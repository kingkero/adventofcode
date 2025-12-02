package day02

import "testing"

func TestIsInvalidIDPart01(t *testing.T) {
	examples := map[string]bool{
		"101":        false,
		"11":         true,
		"12":         false,
		"1111":       true,
		"1212":       true,
		"1112":       false,
		"1188511885": true,
		"1188511886": false,
		"222220":     false,
		"222221":     false,
		"222223":     false,
		"222224":     false,
	}

	for id, expectedResult := range examples {
		result := isInvalidIDPart01(id)
		if result != expectedResult {
			t.Errorf("%s: expected %v, got %v", id, expectedResult, result)
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

func TestIsInvalidIDPart02(t *testing.T) {
	examples := map[string]bool{
		"999":        true,
		"565656":     true,
		"824824824":  true,
		"2121212121": true,
	}

	for id, expectedResult := range examples {
		result := isInvalidIDPart02(id)
		if result != expectedResult {
			t.Errorf("%s: expected %v, got %v", id, expectedResult, result)
		}
	}
}

func TestPart02(t *testing.T) {
	result := Part02([]string{
		"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
	})

	if result != "4174379265" {
		t.Errorf("got %s, want 4174379265", result)
	}
}

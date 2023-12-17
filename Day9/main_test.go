package main

import "testing"

func TestCountUniqueCharacter(t *testing.T) {

	s := "Hello123"

	result := CountUniqueCharacter(s)

	expected := map[string]int{
		"h": 1,
		"e": 1,
		"l": 2,
		"o": 1,
		"1": 1,
		"2": 1,
		"3": 1,
	}

	for char, count := range expected {
		if result[char] != count {
			t.Errorf("expected count of '%s' is %d but got : %d", char, count, result[char])
		}
	}

}

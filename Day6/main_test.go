package main

import "testing"

func TestWordCount(t *testing.T) {

	result := WordCount("This is test. this test is simple")

	expected := map[string]int{
		"this" : 2,
		"is" : 2,
		"test" : 2,
		"simple" : 1,
	}
	
	for word, count := range expected {
		if result[word] != count {
			t.Errorf("Expected %s words count shoul be %d but got %d", word, count, result[word])
		}
	}
}
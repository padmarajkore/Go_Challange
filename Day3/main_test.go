package main

import "testing"

func TestFindMax(t *testing.T) {
	Number := []int{4,5,7,8,9,6}

	result:= FindMax(Number)

	if result != 9 {
		t.Error("Expected Max value 9, got:", result)
	}

	//test case for empty slice
	EmptySlice := []int{}
	result = FindMax(EmptySlice)

	if result != -1 {
		t.Error("Expected error for empty slice, got:", result)
	}

}
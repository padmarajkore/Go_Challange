package main

import "testing"

func TestProductNonNegative(t *testing.T) {
	numbers := []int{7,8,-9,8,6,-7,8,9,7}
	result := ProductNonNegative(numbers)

	if result != 1354752 {
		t.Error("expected product to be 1354752, but found:", result)
	}

	numbers = []int{}
	result = ProductNonNegative(numbers)

	if result != 1 {
		t.Error("expected product to be 1, but found:", result)
	}

	numbers = []int{-9,-8,-9,-5}
	result = ProductNonNegative(numbers)

	if result != 1 {
		t.Error("expected product to be 1, but found:", result)
	}
} 
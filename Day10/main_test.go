package main

import (
	"reflect"
	"testing"
)

func TestUniqueElements(t *testing.T) {

	input := []int{2, 4, 5, 6, 7, 4, 5}

	result := UniqueElements(input)

	expected := []int{2, 4, 5, 6, 7}

	if !reflect.DeepEqual(result, expected) {
		t.Error("expected result is ", expected, "but got :", result)
	}
}

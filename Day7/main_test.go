package main

import "testing"

func TestPositiveSum(t *testing.T) {

	result := PositiveSum([]int{7,-8,9,-6,7,-5,-5,4,-5,7,-3,2,-1})

	if result != 36{
		t.Error("Expected result should be 36 and got:",result)
	}

	result = PositiveSum([]int{6,7,-9,0,7,-8,-7})

	if result != 20{
		t.Error("Expected result should be 20 and got:",result)
	}

}
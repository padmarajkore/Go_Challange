package main

import "testing"

func TestIsLeap(t *testing.T) {

	if !IsLeap(2000){
		t.Error("Expected 2024 is a leap here")
	}

	if !IsLeap(2024){
		t.Error("Expected 2024 is a leap here")
	}

	if  IsLeap(1900){
		t.Error("Expected 2024 is a leap here")
	}

	if  IsLeap(2022){
		t.Error("Expected 2024 is a leap here")
	}
}
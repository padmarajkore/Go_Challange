package main

import "testing"

func TestIsPrime(t* testing.T) {

	if !IsPrime(2) {
		t.Error("Expected 2 to be prime")
	}

	if !IsPrime(17) {
		t.Error("Expected 17 to be prime")
	}

	if IsPrime(4) {
		t.Error("Expected 4 not to be prime")
	}

	if IsPrime(10) {
		t.Error("Expected 10 not to be prime")
	}

	if !IsPrime(13) {
		t.Error("Expected 13 to be prime")
	}

	if IsPrime(1) {
		t.Error("Expected 1 not to be prime")
	}
}
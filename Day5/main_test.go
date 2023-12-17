package main

import "testing"

func TestFacotrial (t *testing.T) {

	result := Facotrial(3)

	if result !=6 {
		t.Error("Expected 6 as a factorial of 3 but got :", result)
	}

	result = Facotrial (0)

	if result != 1 {
		t.Error("Expected 1 as a facotial of 0 but got :", result)
	}
}
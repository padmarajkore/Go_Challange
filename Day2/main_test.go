package main

import "testing"

func TestIsPalindrome(t* testing.T) {

	if !IsPalindrome("racecar"){
		t.Error("Expected 'racecar' to be a palindrome")
	}	

	if !IsPalindrome("A man, a plan, a canal, Panama"){
		t.Error("Expected 'A man, a plan, a canal, Panama' to be palindrome")
	}

	if IsPalindrome("hello") {
		t.Error("Expected 'hello' not to be palindrome")
	}

	if IsPalindrome("not a palindrome") {
		t.Error("Expected 'not a palindrome' not to be a palindrome")
	}
}

package main

import (
	"strings"
	"unicode"
	"fmt"
)

func IsPalindrome(s string) bool{
	s = preprocessString(s)

	for i, j := 0, len(s)-1; i < j; i , j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true

}

func preprocessString(s string) string {
	var sb strings.Builder
	
	for _, char := range s {
		if unicode.IsLetter(char) {
			sb.WriteRune(unicode.ToLower(char))
		}
	}

	return sb.String()
}

func main() {

	fmt.Println(IsPalindrome("A man, a plan, a canal, panama"))
	fmt.Println(IsPalindrome("racecar"))
	fmt.Println(IsPalindrome("hello"))
	fmt.Println(IsPalindrome("not a palindrome"))
}
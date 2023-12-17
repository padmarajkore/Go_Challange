package main

import (
	"strings"
	"unicode"
)

func CountUniqueCharacter(s string) map[string]int {

	charCount := make(map[string]int)

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			charLower := strings.ToLower(string(char))
			charCount[charLower]++
		}
	}
	return charCount
}

func main() {
	return
}

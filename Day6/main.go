package main

import (
	"regexp"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)

	words := regexp.MustCompile(`\b\w+\b`).FindAllString(s, -1)

	for _, word := range words {

		word = strings.ToLower(word)
		wordCount[word]++
	}

	return wordCount
}
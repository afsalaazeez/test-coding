package myfunc

import (
	"strings"
)

// CountWords parses a string and return the number of words in the passage
func CountWords(passage string) int {
	words := strings.Fields(passage)
	unique := make(map[string]int)
	for _, word := range words {
		unique[word]++
	}
	return len(unique)
}

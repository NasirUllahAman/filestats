package utils

import "strings"

func CountCharacters(data string) int {
	count := 0
	return count
}
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}

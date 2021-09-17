package utils

import "strings"

func CharCount(data []byte) int {
	return len(data)
}
func WordCount(data []byte) int {
	count := 0
	for i := 0; i < len(data); i++ {
		if string(data[i]) == " " || i == len(data)-1 {
			count++
		}
	}
	return count
}

func LineCount(data []byte) int {
	count := 0
	for i := 0; i < len(data); i++ {
		if string(data[i]) == "\n" || i == len(data)-1 {
			count++
		}
	}
	return count
}

func DuplicateWord(data string) map[string]int {

	input := strings.Fields(data)
	count := make(map[string]int)
	for _, word := range input {
		_, matched := count[word]
		if matched {
			count[word] += 1
		} else {
			count[word] = 1
		}
	}
	return count
}

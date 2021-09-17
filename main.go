package main

import (
	"fmt"

	"github.com/NasirUllahAman/filestats/utils"
)

func main() {
	filePath := "lorem.txt"
	data, err := utils.ReadLocalFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(data))
	charCount := utils.CharCount(data)
	wordCount := utils.WordCount(data)
	LineCount := utils.LineCount(data)
	DuplicateWord := utils.DuplicateWord(string(data))

	fmt.Printf("Word Count: %d\nCharacter Count: %d\nline count:%d\n, DuplicateWord:\n", wordCount, charCount, LineCount, DuplicateWord)
}

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
	fmt.Printf("Word Count: %d\nCharacter Count: %d\n line count:%d \n", wordCount, charCount, LineCount)
}

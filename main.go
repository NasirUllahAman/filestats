package main

import (
	"fmt"

	"github.com/NasirUllahAman/filestats/utils"
)

func main() {
	filePath := "Readme.md"
	data, err := utils.ReadFile(filePath)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	utils.CountCharacters("the main vsc page have")
	fmt.Println(data)
}

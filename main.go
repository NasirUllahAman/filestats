package main

import (
	"fmt"
	"io/ioutil"

	"github.com/NasirUllahAman/filestats/utils"
)

func main() {
	filePath := "Readme.md"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(data))
	utils.CharCount(data)
	fmt.Println()
	fmt.Println()

}

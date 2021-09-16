//package utils
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/NasirUllahAman/filestats/utils"
)

func ReadFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
func main() {
	filePath := "Readme.md"
	data, err := utils.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(data)
	ReadFile(data)

}

package utils

import (
	"io/ioutil"
)

func ReadLocalFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

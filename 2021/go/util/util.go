package util

import (
	"fmt"
	"io/ioutil"
)

func ReadDay(day int) (string, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("inputs/%d.txt", day))
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ReadExample(day int) (string, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("../inputs/%d.example.txt", day))
	if err != nil {
		return "", err
	}
	return string(content), nil
}

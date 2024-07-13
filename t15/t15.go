package main

import (
	"errors"
	"fmt"
)

// createHugeString создает большую строку заданного размера
func createHugeString(size int) string {
	hugeString := make([]rune, size)
	for i := range hugeString {
		hugeString[i] = 'a'
	}
	return string(hugeString)
}

var justString string

func someFunc() error {
	v := createHugeString(1 << 10)
	if len(v) < 100 {
		return errors.New("строка слишком короткая")
	}
	justString = string([]rune(v[:100]))
	return nil
}

func main() {
	err := someFunc()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(justString)
	}
}

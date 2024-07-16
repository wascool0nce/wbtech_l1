package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	s = strings.ToLower(s)

	charMap := make(map[rune]bool)
	for _, char := range s {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
}

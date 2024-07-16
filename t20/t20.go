package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	n := len(words)
	for i := 0; i < n/2; i++ {
		words[i], words[n-1-i] = words[n-1-i], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	words := "snow dog sun"
	sdrow := reverseWords(words)
	fmt.Println(sdrow)
}

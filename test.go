package main

import (
	"fmt"
	"strconv"
)

func ValidISBN10(isbn string) bool {
	count := 0
	if len(isbn) < 9 {
	  return false
	}
	
	for idx, char := range isbn {
	  num, err := strconv.Atoi(string(char))
	  if err != nil && idx < 9 {
		  return false
	  }
	  count += (idx+1)*num
	  if idx < 9 && num > 10 {
		return false
	  }
	  if idx == 9 && string(char) == "X" {
		return true
	  }
	}
	if count % 11 != 0 {
	  return false
	}
	return true // implement proper solution here
  }


func main() {
	ValidISBN10("1112223339")
}

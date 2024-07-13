package main

import "fmt"

func binarySearch(needle int, haystack []int) bool {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(binarySearch(3, arr))
}

package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]

	var less, greater []int
	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}

func main() {
	arr := []int{1, 3, 4, 6, 2, 7, 5, 8, 10, 9}
	sortArr := quickSort(arr)
	fmt.Println(sortArr)
}

package main

import "fmt"

func remove(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("i меньше 0 либо больше размера слайса")
		return nil
	}
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный слайс:", slice)

	i := 2
	slice = remove(slice, i)

	fmt.Println("Слайс после удаления элемента:", slice)
}

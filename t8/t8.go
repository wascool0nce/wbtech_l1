package main

import (
	"fmt"
)

func setBit(num *int64, i uint) {
	*num |= 1 << i
}

func main() {
	var i uint = 1
	var num int64 = 100
	for ; i < 10; i++ {
		// Пример исходного числа
		fmt.Printf("Исходное значение: %d\n", num)

		// Устанавливаем 3-й бит в 1
		setBit(&num, i)
		fmt.Printf("После установки i бита: %d\n", num)

	}
}

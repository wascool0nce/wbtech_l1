package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 10
	fmt.Println("Значение а:", a, "Значение b:", b)
	a, b = b, a
	fmt.Println("Значение а:", a, "Значение b:", b)
}
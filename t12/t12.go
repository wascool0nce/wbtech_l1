package main

import (
	"fmt"
)

func main() {
	animals := []string{"cat", "cat", "dog", "cat", "tree"}

	countAnimals := make(map[string]int)

	for _, a := range animals {
		countAnimals[a]++
	}

	for a, num := range countAnimals {
		fmt.Printf("%s встречается %d раз(а) \n", a, num)
	}
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	numbers := [5]int{2, 4, 6, 8, 10}

	// Увеличиваем счетчик WaitGroup на 5

	for i := 0; i < len(numbers); i++ {
		wg.Add(1) 
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			fmt.Println(numbers[i] * numbers[i])
			mu.Unlock()
		}(i) // Передаем i в качестве параметра
	}

	wg.Wait() // Ждем завершения всех горутин
}

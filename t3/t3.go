package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
func main() {
	var sum int
	var wg sync.WaitGroup
	var mu sync.Mutex
	numbers := [5]int{2, 4, 6, 8, 10}

	for i := 0; i < len(numbers); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			sum = sum + numbers[i]*numbers[i]
			mu.Unlock()
		}(i) // Передаем i в качестве параметра
	}

	wg.Wait() // Ждем завершения всех горутин
	fmt.Println(sum)
}

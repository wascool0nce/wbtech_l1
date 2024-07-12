package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Убираем мьютекс
	counter := 0
	goroutines := 100000000

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			// Увеличиваем счетчик без синхронизации
			counter++
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}

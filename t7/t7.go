package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu   sync.RWMutex
		data = make(map[string]int)
		wg   sync.WaitGroup
	)

	numReaders := 5
	numWriters := 5
	wg.Add(numReaders + numWriters)

	for i := 0; i < numReaders; i++ {
		go func(id int) {
			defer wg.Done()

			mu.RLock()
			defer mu.RUnlock()

			key := fmt.Sprintf("key%d", id%numWriters) // Чтобы читать уже записанные данные
			value := data[key]
			fmt.Printf("Reader %d прочитал из мапы: %s -> %d\n", id, key, value)
		}(i)
	}

	for i := 0; i < numWriters; i++ {
		go func(id int) {
			defer wg.Done()

			// Защищаем доступ к map для записи с помощью Lock()
			mu.Lock()
			defer mu.Unlock()

			// Обновляем значения в map
			key := fmt.Sprintf("key%d", id)
			data[key] = id * 10
			fmt.Printf("Writer %d обновил мапу с ключем %s\n", id, key)
		}(i)
	}

	wg.Wait()

	// Выводим все элементы map после завершения работы всех горутин
	fmt.Println("Финальная мапа:", data)
}

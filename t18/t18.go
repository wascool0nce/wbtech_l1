package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	goroutines := 1000
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}

	wg.Wait()

	fmt.Printf("Итоговое значение счетчика: %d\n", counter.Value())
}

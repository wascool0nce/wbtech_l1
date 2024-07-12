package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	stop     bool
	stopLock sync.Mutex
)

func workerStopChan(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Воркер остановлен!")
			return
		default:
			fmt.Println("Воркер работает")
			time.Sleep(1 * time.Second)
		}
	}
}

func workerStopCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Воркер остановлен!")
			return
		default:
			fmt.Println("Воркер работает")
			time.Sleep(1 * time.Second)
		}
	}
}

func workerStopWG(wg *sync.WaitGroup, stopChan chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-stopChan:
			fmt.Println("Воркер остановлен!")
			return
		default:
			fmt.Println("Воркер работает")
			time.Sleep(1 * time.Second)
		}
	}
}

func workerStopMutex() {
	for {
		stopLock.Lock()
		if stop {
			stopLock.Unlock()
			fmt.Println("Воркер остановлен!")
			return
		}
		stopLock.Unlock()
		fmt.Println("Воркер работает")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Используем канал для сигнала остановки горутины
	stopChan := make(chan struct{})
	go workerStopChan(stopChan)

	time.Sleep(3 * time.Second)
	close(stopChan)

	time.Sleep(3 * time.Second)

	fmt.Println("<----------------------------------------->")

	// Используем контекст для остановки горутины
	ctx, cancel := context.WithCancel(context.Background())
	go workerStopCtx(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)

	fmt.Println("<----------------------------------------->")

	// Используем флаг с мьютексом чтобы остановить горутину
	go workerStopMutex()
	time.Sleep(3 * time.Second)
	stopLock.Lock()
	stop = true
	stopLock.Unlock()
	time.Sleep(3 * time.Second)

	fmt.Println("<----------------------------------------->")

	var wg sync.WaitGroup
	stopChanWG := make(chan struct{})
	wg.Add(1)
	go workerStopWG(&wg, stopChanWG)

	time.Sleep(6 * time.Second)
	close(stopChanWG)
	wg.Wait()
}

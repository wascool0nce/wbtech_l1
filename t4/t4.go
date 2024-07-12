package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func writeCh(in chan<- int, ctx context.Context) {
	for i := 0; true; i++ {
		select {
		case <-ctx.Done():
			close(in)
			return
		default:
			in <- i
			time.Sleep(1 * time.Second)
		}

	}
}
// В функции worker добавлен параметр wg *sync.WaitGroup, 
// который используется для отслеживания активных воркеров. 
// После завершения работы каждый воркер вызывает wg.Done().
// В главной функции main после получения сигнала 
// о завершении (cancel() вызывается при получении Ctrl+C), 
// ожидается завершение всех воркеров с помощью wg.Wait().
func workers(numWorker int, wg *sync.WaitGroup, outs <-chan int) {
	defer wg.Done()
	for out := range outs {
		fmt.Printf("Воркер-%d выполнил: %d\n", numWorker, out)
	}
	fmt.Printf("Воркер-%d остановился\n", numWorker)
}

func main() {

	numWorkers := flag.Int("workers", 4, "number of workers")
	flag.Parse()

	chData := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	var wg sync.WaitGroup
	for i := 1; i <= *numWorkers; i++ {
		wg.Add(1)
		go workers(i, &wg, chData)
	}

	go writeCh(chData, ctx)
	select {
	case <-sigCh:
		fmt.Println("\nЗавершение работы...")
		cancel()
	}

	wg.Wait()
	fmt.Println("Все воркеры вышли. Программа выключена")
}

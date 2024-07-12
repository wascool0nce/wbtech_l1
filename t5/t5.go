package main

import (
	"flag"
	"fmt"
	"time"
)

func writeData(in chan<- int) {
	for i := 1; true; i++ {
		in <- i
		fmt.Println("Значение i:",i)
		time.Sleep(1 * time.Second)
	}
}

func readData(out <-chan int) {
	for {
		value := <-out
		fmt.Println("Значение из канала:", value)
	}
}

func main() {
	timeWorks := flag.Int("time", 10, "time of work")
	flag.Parse()

	chData := make(chan int)
	go writeData(chData)

	go readData(chData)


	time.Sleep(time.Duration(*timeWorks) * time.Second)

	fmt.Println("Программа завершена.")
}

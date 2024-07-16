package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Начало сна")
	sleep(5 * time.Second)
	fmt.Println("Конец сна")
}

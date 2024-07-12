package main

import (
    "fmt"
)

func main() {
    // Массив чисел
    numbers := []int{1, 2, 3, 4, 5}

    // Создаем каналы
    ch1 := make(chan int)
    ch2 := make(chan int)

    // Горутина для записи чисел в первый канал
    go func() {
        for _, num := range numbers {
            ch1 <- num
        }
        close(ch1) // Закрываем канал после записи всех чисел
    }()

    // Горутина для чтения из первого канала, умножения на 2 и записи во второй канал
    go func() {
        for num := range ch1 {
            ch2 <- num * 2
        }
        close(ch2) // Закрываем канал после обработки всех чисел
    }()

    // Чтение из второго канала и вывод в stdout
    for result := range ch2 {
        fmt.Println(result)
    }
}
package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	goroutineAndWaitGroup(numbers)
	fmt.Println("--------")
	goroutineAndUnbufferedChannel(numbers)
	fmt.Println("--------")
	goroutineAndBufferedChannel(numbers)
	fmt.Println("--------")
}

func goroutineAndWaitGroup(numbers []int) {
	var wg sync.WaitGroup // создание WaitGroup для ожидания завершения горутин

	for _, number := range numbers {
		wg.Add(1) // добавление счетчика для каждой горутины
		go func(number int) {
			defer wg.Done() // уменьшение счетчика для каждой горутины
			square := number * number
			fmt.Println(square)
		}(number)
		wg.Wait() // ожидание завершения работы всех горутин
	}
}

func goroutineAndUnbufferedChannel(numbers []int) {
	result := make(chan int) // создание канала

	for _, number := range numbers {
		go func(number int) {
			square := number * number
			result <- square // отправление результата в канал
		}(number)
	}
	for range numbers {
		fmt.Println(<-result) // получение результата из канала
	}
	close(result) // закрытие канала
}

func goroutineAndBufferedChannel(numbers []int) {
	result := make(chan int, len(numbers)) // создание канала размером длины слайса numbers
	// буфферизированный используется для асинхронной работы с данными, он не блокируется пока не будет заполнен/опустошен

	for _, number := range numbers {
		go func(number int) {
			square := number * number
			result <- square // отправление результата в канал
		}(number)
	}
	for range numbers {
		fmt.Println(<-result) // получение результата из канала
	}
	close(result) // закрытие канала
}

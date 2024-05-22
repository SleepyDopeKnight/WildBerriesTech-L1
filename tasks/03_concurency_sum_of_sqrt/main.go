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
	summary := 0

	var wg sync.WaitGroup // создание WaitGroup для ожидания завершения горутин

	for _, number := range numbers {
		wg.Add(1) // добавление счетчика для каждой горутины

		go func(number int) {
			defer wg.Done() // уменьшение счетчика для каждой горутины

			square := number * number
			summary += square
		}(number)
		wg.Wait() // ожидание завершения работы всех горутин
	}

	fmt.Println(summary)
}

func goroutineAndUnbufferedChannel(numbers []int) {
	result := make(chan int) // создание канала
	summary := 0

	for _, number := range numbers {
		go func(number int) {
			result <- number * number // отправление результата в канал
		}(number)
	}

	for range numbers {
		summary += <-result // получение результата из канала и сложение полученных его значений
	}

	fmt.Println(summary)
	close(result) // закрытие канала
}

func goroutineAndBufferedChannel(numbers []int) {
	result := make(chan int, len(numbers)) // создание канала
	summary := 0

	for _, number := range numbers {
		go func(number int) {
			result <- number * number // отправление результата в канал
		}(number)
	}

	for range numbers {
		summary += <-result // получение результата из канала и сложение полученных его значений
	}

	fmt.Println(summary)
	close(result) // закрытие канала
}

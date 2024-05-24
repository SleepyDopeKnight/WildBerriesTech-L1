package main

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	chan1 := make(chan int)
	chan2 := make(chan int)

	go writeInFirstChan(numbers, chan1) // запускаем горутину для записи слайса в 1ый канал

	go mul2FromFirstChanInSecond(chan1, chan2) // запускаем горутину для записи чисел *2 во 2ой канал

	for number := range chan2 { // считываем числа из второго канала
		fmt.Println(number)
	}
}

func writeInFirstChan(numbers []int, first chan int) {
	for _, number := range numbers {
		first <- number // записываем числа из слайса в канал
	}

	close(first) // закрываем канал
}

func mul2FromFirstChanInSecond(first chan int, second chan int) {
	for number := range first {
		second <- number * 2 // умножаем числа на 2 из первого канала и записываем во второй
	}

	close(second) // закрываем канал
}

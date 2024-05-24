package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	stopSignal, cancel := context.WithTimeout(context.Background(), time.Second*10) // создает контекст, который закончит программу
	defer cancel()                                                                  // очищаем созданный контекст
	start(stopSignal)
}

func start(stopSignal context.Context) {
	var wg sync.WaitGroup

	data := make(chan int) // создаем канал, в который будем записывать данные

	wg.Add(2)

	go generateData(stopSignal, data, &wg) // запускаем горутину, в которую передаем канал для записи данных и контекст, для завершения работы рутины

	go worker(stopSignal, data, &wg) // запуск работника
	defer close(data)                // закрываем канал
	wg.Wait()                        // ожидание завершения работы всех горутин
}

func worker(ctx context.Context, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшение счетчика для каждой горутины

	for {
		select {
		case <-ctx.Done(): // если пришел контекст выхода из программы, то завершаем программу
			fmt.Println("Worker exiting")
			return
		case number, ok := <-data: // записываем сгенерированные данные в канал
			if !ok { // проверяем, не закрылся ли канал
				fmt.Println("Channel is closed")
				return
			}

			fmt.Println(number)
		}
	}
}

func generateData(ctx context.Context, data chan<- int, wg *sync.WaitGroup) {
	wg.Done()

	for {
		select {
		case <-ctx.Done(): // если пришел контекст выхода из программы, то завершаем программу
			fmt.Println("Stop generating")
			return
		default:
			time.Sleep(time.Second)
			data <- rand.Int() // записываем сгенерированные данные в канал для чтения работниками
		}
	}
}

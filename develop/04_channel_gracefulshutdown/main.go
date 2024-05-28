package main

/*Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые
читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var workers int

	stopSignal, cancel := context.WithCancel(context.Background()) // создает контекст, с помощью которого можно выйти из программы
	defer cancel()                                                 // для отмены созданого контекста
	gracefulShutdown(cancel)                                       // выход из программы на ctlc+c
	fmt.Println("Count of workers:")

	if _, err := fmt.Fscan(os.Stdin, &workers); err != nil || workers <= 0 {
		fmt.Println("Invalid number of workers") // считываем из stdin количество работников
		return
	}

	start(stopSignal, workers)
}

func start(stopSignal context.Context, workers int) {
	var wg sync.WaitGroup

	data := make(chan int)            // создаем канал, в который будем записывать данные
	go generateData(stopSignal, data) // запускаем горутину, в которую передаем канал для записи данных и контекст, для завершения работы рутины

	for i := 0; i < workers; i++ {
		wg.Add(1) // добавление счетчика для каждой горутины

		go worker(stopSignal, data, &wg) // запуск работников
	}

	defer close(data) // закрываем канал
	wg.Wait()         // ожидание завершения работы всех горутин
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

func generateData(ctx context.Context, data chan<- int) {
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

func gracefulShutdown(cancel context.CancelFunc) {
	go func() {
		signalChan := make(chan os.Signal, 1)   // буфферизированный канал(чтобы не заблокировался), принмающий сигнала системы
		signal.Notify(signalChan, os.Interrupt) // оповещает канал о сигнале, полученный от os.Interrupt(подписывает канал на получения сигнала)
		defer signal.Stop(signalChan)

		<-signalChan // ожидаю сигнала стоп от канала
		cancel()     // очищаем созданный контекст
	}()
}

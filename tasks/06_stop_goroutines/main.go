package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	exitWithContext()
	fmt.Println("--------")
	exitWithChannel()
	fmt.Println("--------")
	exitWithTimer()
	fmt.Println("--------")
	exitWithAtomicVar()
	fmt.Println("--------")
}

func exitWithContext() {
	ctx, cancel := context.WithCancel(context.Background()) // создает контекст, с помощью которого можно выйти из программы
	var wg sync.WaitGroup

	wg.Add(1)                                          // добавление счетчика для каждой горутины
	go func(ctx context.Context, wg *sync.WaitGroup) { // запуск горутины
		defer wg.Done() // уменьшение счетчика для каждой горутины
		for {
			select {
			case <-ctx.Done(): // получаем завершение работы контекста
				fmt.Println("Exit with context")
				return
			default:
				time.Sleep(time.Millisecond * 500)
				fmt.Println("Goroutine is working")
			}
		}
	}(ctx, &wg)
	time.Sleep(time.Second * 2) // даем рутине поработать 2 секунды перед выходом
	cancel()                    // для отмены созданого контекста
	wg.Wait()                   // ожидание завершения работы всех горутин
}

func exitWithChannel() {
	semaphore := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)                                              // добавление счетчика для каждой горутины
	go func(semaphore chan struct{}, wg *sync.WaitGroup) { // запуск горутины
		defer wg.Done() // уменьшение счетчика для каждой горутины
		for {
			select {
			case <-semaphore: // получаем сигнал, что канал закрылся
				fmt.Println("Exit with closed channel")
				return
			default:
				time.Sleep(time.Millisecond * 500)
				fmt.Println("Goroutine is working")
			}
		}
	}(semaphore, &wg)
	time.Sleep(time.Second * 2) // даем рутине поработать 2 секунды перед выходом
	close(semaphore)            // закрываем канал
	wg.Wait()                   // ожидание завершения работы всех горутин
}

func exitWithTimer() { // тоже самое что выход по таймауту контекста
	timer := time.NewTimer(time.Second * 2)
	var wg sync.WaitGroup
	wg.Add(1)                                             // добавление счетчика для каждой горутины
	go func(timer <-chan time.Time, wg *sync.WaitGroup) { // запуск горутины
		defer wg.Done() // уменьшение счетчика для каждой горутины
		for {
			select {
			case <-timer: // получаем сигнал об окончании таймера
				fmt.Println("Exit with timer")
				return
			default:
				time.Sleep(time.Millisecond * 500)
				fmt.Println("Goroutine is working")
			}
		}
	}(timer.C, &wg)
	wg.Wait() // ожидание завершения работы всех горутин
}

func exitWithAtomicVar() {
	var run int32 = 1
	var wg sync.WaitGroup
	wg.Add(1)                     // добавление счетчика для каждой горутины
	go func(wg *sync.WaitGroup) { // запуск горутины
		defer wg.Done() // уменьшение счетчика для каждой горутины
		for atomic.LoadInt32(&run) == 1 {
			time.Sleep(time.Millisecond * 500)
			fmt.Println("Goroutine is working")
		}
		fmt.Println("Exit with atomic")
	}(&wg)
	time.Sleep(time.Second * 2) // даем рутине поработать 2 секунды перед выходом
	atomic.StoreInt32(&run, 0)  // устанавливает атомарный флаг в 0
	wg.Wait()                   // ожидание завершения работы всех горутин
}

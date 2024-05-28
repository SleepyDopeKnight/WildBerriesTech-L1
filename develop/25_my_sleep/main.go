package main

/*
Реализовать собственную функцию sleep.
*/
import (
	"context"
	"fmt"
	"time"
)

func main() {
	sleepWithContext(2 * time.Second)
	sleepWithTimer(2 * time.Second)
	sleepWithDelay(2 * time.Second)
}

func sleepWithContext(duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration) // создает контекст, с помощью которого можно выйти из программы
	defer cancel()                                                     // отменяем созданный контекст

	<-ctx.Done() // по истечению времени приходит сигнал об окончании контекста
	fmt.Printf("Exit after context timeout %s seconds\n", duration)
}

func sleepWithTimer(duration time.Duration) { // тоже самое что выход по таймауту контекста
	timer := time.NewTimer(duration) // создаем таймер

	<-timer.C // по истечению таймера приходит сигнал в канал таймера
	fmt.Printf("Exit after timer timeout %s seconds\n", duration)
}

func sleepWithDelay(duration time.Duration) {
	<-time.After(duration) // ожидаем сигнал в канал по истечению времени
	fmt.Printf("Exit after delay timeout %s seconds\n", duration)
}

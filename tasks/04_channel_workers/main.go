package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

func main() {
	stopSignal, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefulShutdown(cancel)
	fmt.Println("Count of workers:")
	var workers int
	var wg sync.WaitGroup
	fmt.Fscan(os.Stdin, &workers)
	data := make(chan int)
	go generateData(stopSignal, data)
	for i := 1; i < workers; i++ {
		wg.Add(1)
		go worker(stopSignal, data, &wg)
	}
}

func worker(ctx context.Context, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exit from program")
			return
		default:
			fmt.Println(<-data)
		}
	}
}

func generateData(ctx context.Context, data chan<- int) {
	defer close(data)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exit from program")
			return
		default:
			data <- rand.Int()
		}
	}
}

func start(ctx context.Context) {
}

func gracefulShutdown(cancel context.CancelFunc) {
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt)
		defer signal.Stop(signalChan)
		<-signalChan
		cancel()
	}()
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterWithMutex struct {
	mutex       sync.Mutex
	countMutex  int
	countAtomic int32
}

func main() {
	var counterMutex CounterWithMutex

	var wg sync.WaitGroup

	counterMutex.start(&wg)
}

func (c *CounterWithMutex) incrementMutex() { // увеличение счетчика через мьютекс
	c.mutex.Lock() // блокировка других горутин для записи
	c.countMutex++
	c.mutex.Unlock() // разбловкировка рутин
}

func (c *CounterWithMutex) incrementAtomic() { // увлечение счетчика через атомарную переменную
	atomic.AddInt32(&c.countAtomic, 1)
}

func (c *CounterWithMutex) start(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ { // запускаем 10 горутин
		wg.Add(1) // добавление счетчика для каждой горутины

		go func() {
			defer wg.Done() // уменьшение счетчика для каждой горутины
			c.incrementMutex()
			c.incrementAtomic()
		}()
	}
	wg.Wait() // ожидание завершения работы всех горутин
	fmt.Println("Mutex counter", c.countMutex)
	fmt.Println("Atomic counter", c.countAtomic)
}

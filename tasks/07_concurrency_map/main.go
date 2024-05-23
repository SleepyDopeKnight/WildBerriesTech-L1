package main

import (
	"fmt"
	"sync"
)

func main() {
	sliceNumbers := []int{2, 4, 6, 8, 10}
	writeWithMutex(sliceNumbers)
	fmt.Println("--------")
	writeWithSyncMap(sliceNumbers)
	fmt.Println("--------")

}

func writeWithMutex(sliceNumbers []int) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	mapNumbers := make(map[int]int)
	for index, number := range sliceNumbers {
		wg.Add(1) // добавление счетчика для каждой горутины

		go func() {
			defer wg.Done() // уменьшение счетчика для каждой горутины

			mutex.Lock() // блокировка других горутин для записи значений в мапу
			mapNumbers[index] = number
			mutex.Unlock() // разбловкировку рутин
		}()
		wg.Wait() // ожидание завершения работы всех горутин
	}
	fmt.Println(mapNumbers)
}

func writeWithSyncMap(sliceNumbers []int) {
	var wg sync.WaitGroup
	var mapNumbers sync.Map // под капотом уже имеет mutex

	for index, number := range sliceNumbers {
		wg.Add(1) // добавление счетчика для каждой горутины

		go func() {
			defer wg.Done() // уменьшение счетчика для каждой горутины

			mapNumbers.Store(index, number) // сохранение значений в мапу
		}()
		wg.Wait() // ожидание завершения работы всех горутин
	}

	mapNumbers.Range(func(key, value any) bool { // проходимся по всем ключам и значениям
		fmt.Println(key, value)
		return true // если true - продолжает иттерацию, иначе останавливает
	})
}

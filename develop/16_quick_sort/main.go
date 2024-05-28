package main

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/
import "fmt"

func main() {
	numbers := []int{8, 7, 9, 14, 2, 23, 16, 4, 8}
	numbers = quickSort(numbers)
	fmt.Println(numbers)
}

func quickSort(numbers []int) []int {
	if len(numbers) < 2 { // если в массиве 1 элемент - он уже отсортирован
		return numbers
	}

	var less, greater []int // слайсы для значений меньших или больших чем elem

	elem := numbers[0]
	for _, number := range numbers[1:] { // тк elem 0 элемент, начинаем цикл с 1го
		if number <= elem {
			less = append(less, number)
		} else {
			greater = append(greater, number)
		}
	}

	result := append(quickSort(less), elem)        // рекурсивная сортировка слайса less
	result = append(result, quickSort(greater)...) // рекурсивная сортировка слайса greater и объединение в результат

	return result
}

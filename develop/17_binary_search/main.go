package main

/*
Реализовать бинарный поиск встроенными методами языка.
*/
import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10, 11}
	number := binarySearch(numbers, 4)
	fmt.Println(number)
	number = binarySearch(numbers, 10)
	fmt.Println(number)
	number = binarySearch(numbers, 9)
	fmt.Println(number)
}

func binarySearch(numbers []int, searchedNumber int) int {
	left, right := 0, len(numbers)-1 // определяем границы поиска числа

	for left <= right {
		middle := (left + right) / 2 // находим индекс середины слайса

		switch {
		case searchedNumber < numbers[middle]:
			right = middle - 1 // если искомое число меньше середины, то мы ищем в левой части слайса
		case searchedNumber > numbers[middle]:
			left = middle + 1 // если искомое число больше середины, то мы ищем в правой части слайса
		default:
			return middle // возвращаем найденное значение
		}
	}

	return -1 // возвращаем -1, если нет искомого элемента
}

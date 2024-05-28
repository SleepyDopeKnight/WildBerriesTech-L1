package main

/*
Удалить i-ый элемент из слайса.
*/
import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println(numbers) // 2 4 6 8 10
	numbers = deleteWithAppend(numbers, 2)
	fmt.Println(numbers) // 2 4 8 10
	numbers = deleteWithCopy(numbers, 3)
	fmt.Println(numbers) // 2 4 8
}

func deleteWithAppend(numbers []int, i int) []int {
	return append(numbers[:i], numbers[i+1:]...)
} // разбиваем слайс на 2 части: до элемента i и после элемента i и объединяем обратно обратно с помощью append

func deleteWithCopy(numbers []int, i int) []int {
	copy(numbers[i:], numbers[i+1:])   // копируем элементы от i+1 на место i, сдвигая слайс влево
	numbers = numbers[:len(numbers)-1] // сокращаем длину, чтобы убрать дубликат в самом конце слайса

	return numbers
} // разбиваем слайс на 2 части: до элемента i и после элемента i и объединяем копирую без элемента i

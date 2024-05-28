package main

/*
Реализовать пересечение двух неупорядоченных множеств.
*/
import "fmt"

func main() {
	var commonData []string // слайс для пересечений

	// заполняем можества
	set1 := initSet("banana", "apple", "pineapple", "orange", "cherry")
	set2 := initSet("lemon", "banana", "orange", "kiwi")

	for elem := range set1 { // проходимся по каждому элементу первого множества
		if _, contains := set2[elem]; contains { // проверяем имеет ли второе подмножество такой ключ
			commonData = append(commonData, elem) // записываем пересечения в слайс
		}
	}

	fmt.Println(commonData)
}

func initSet(keys ...string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, key := range keys { // проходимся по слайсу ключей
		set[key] = struct{}{} // записываем ключ в множество
	}

	return set
}

package main

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
import "fmt"

func main() {
	consistency := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, elem := range consistency { // проходимся по каждому элементу последовательности
		set[elem] = struct{}{} // записываем элемент в множество
	}

	fmt.Println(set)
}

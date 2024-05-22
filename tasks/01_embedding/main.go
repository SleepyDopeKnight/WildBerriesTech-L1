package main

import "fmt"

type Human struct {
	Age int
}

func (h *Human) printAge() {
	fmt.Println("Age now: ", h.Age) // вывод переменной Age структуры Human
}

type Action struct {
	Human // встраивание
}

func (a *Action) nextYearAge() {
	a.Age++                                  // инкрементирование переменной Age структуры Action, унаследованной от Human
	fmt.Println("Age on next year: ", a.Age) // вывод переменной Age структуры Action, унаследованной от Human
}

func main() {
	age := 24
	a := Action{Human{Age: age}}
	a.printAge()    // вызов родительского метода структуры Human
	a.nextYearAge() // вызов метода структуры Action
}

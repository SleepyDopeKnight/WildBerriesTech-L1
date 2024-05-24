package main

import "fmt"

func main() {
	variable1 := 1
	variable2 := "string"
	variable3 := true
	variable4 := make(chan int)
	variable5 := make(chan string)
	variable6 := make(chan bool)
	variable7 := 13.37
	typeIdentifier(variable1, variable2, variable3, variable4, variable5, variable6, variable7)
}

func typeIdentifier(variables ...interface{}) {
	for _, variable := range variables {
		switch expr := variable.(type) {
		case int:
			fmt.Println("variable is int", expr)
		case string:
			fmt.Println("variable is string", expr)
		case bool:
			fmt.Println("variable is bool", expr)
		case chan int:
			fmt.Println("variable is chan int", expr)
		case chan string:
			fmt.Println("variable is chan string", expr)
		case chan bool:
			fmt.Println("variable is chan bool", expr)
		default:
			fmt.Println("invalid type", expr)
		}
	}
}

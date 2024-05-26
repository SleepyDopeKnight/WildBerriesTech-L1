package main

import "fmt"

func main() {
	original := "главрыба, abobus"
	reversed := reverseString(original)
	fmt.Println(reversed)
}

func reverseString(str string) string {
	reversed := []rune(str) // переводим строку в слайс рун
	strLen := len(reversed)

	for i, j := 0, strLen-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i] // меняем местами руны
	}

	return string(reversed) // возвращаем слайс рун преобразованный обратно в строку
}

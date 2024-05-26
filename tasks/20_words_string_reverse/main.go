package main

import (
	"fmt"
	"strings"
)

func main() {
	original := "snow dog sun"
	reversed := reverseWordsInString(original)
	fmt.Println(reversed)
}

func reverseWordsInString(str string) string {
	words := strings.Fields(str) // разбиваем строку на слайс слов
	strLen := len(words)

	for i, j := 0, strLen-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i] // меняем местами слова
	}

	return strings.Join(words, " ") // склеиваем после обратно слайс в строку
}

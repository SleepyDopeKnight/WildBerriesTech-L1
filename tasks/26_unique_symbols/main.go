package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	string1 := "abcd"
	string2 := "abCdefAaf"
	string3 := "aabcd"

	fmt.Println(ValidationUniqSymbolsWithSort(string1))
	fmt.Println(ValidationUniqSymbolsWithSort(string2))
	fmt.Println(ValidationUniqSymbolsWithSort(string3))
	fmt.Println("--------")
	fmt.Println(ValidationUniqSymbolsWithMap(string1))
	fmt.Println(ValidationUniqSymbolsWithMap(string2))
	fmt.Println(ValidationUniqSymbolsWithMap(string3))
	fmt.Println("--------")
}

func ValidationUniqSymbolsWithSort(str string) bool {
	str = strings.ToLower(str)         // приводим все символы к одному регистру
	sliceStr := strings.Split(str, "") // разбиваем строку на слайс
	sort.Strings(sliceStr)             // сортируем

	for i := 1; i < len(sliceStr); i++ {
		if sliceStr[i] == sliceStr[i-1] { // по индексу проверяем, если символ равен предидущему - выходим
			return false
		}
	}

	return true
}

func ValidationUniqSymbolsWithMap(str string) bool {
	str = strings.ToLower(str)    // приводим все символы к одному регистру
	mapStr := make(map[rune]bool) // создаем сет, в которые будем записывать символы

	for _, char := range str {
		if mapStr[char] {
			return false
		}

		mapStr[char] = true
	}

	return true
}

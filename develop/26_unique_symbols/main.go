package main

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/
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
	fmt.Println(ValidationUniqSymbolsWithSet(string1))
	fmt.Println(ValidationUniqSymbolsWithSet(string2))
	fmt.Println(ValidationUniqSymbolsWithSet(string3))
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

func ValidationUniqSymbolsWithSet(str string) bool {
	str = strings.ToLower(str)        // приводим все символы к одному регистру
	setStr := make(map[rune]struct{}) // создаем сет, в которые будем записывать символы

	for _, char := range str {
		if _, ok := setStr[char]; ok {
			return false
		}

		setStr[char] = struct{}{}
	}

	return true
}

package main

import "fmt"

func createHugeString(stringSize int) string {
	byteArray := make([]byte, stringSize) // создание слайса байтов заданного размера
	for i := 0; i < stringSize; i++ {
		byteArray[i] = 'z' // заполнение слайса байтов
	}

	return string(byteArray) // приведение массива слайсов в строку
}

// var justString string
// func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
// } если someFunc вызывается в нескольких потоках - может привести к гонке данных, перезаписи и непредсказуемому поведению в justString

// justString[:len] если так делать, мы будем использовать ссылку на весь изначальный массив, он не сможет очиститься,
// пока на него хоть кто-то ссылается и они долго живут
func someFunc(justString string, length int) string { // функция для создания новой подстроки
	return string([]rune(justString)[:length]) // так мы берем len символы изначальной строки и получаем новую строку
}

func main() {
	justString := createHugeString(1 << 10)
	newString := someFunc(justString, 5)
	fmt.Println(newString)
}

package main

import "fmt"

func main() {
	var number int64 = 5 // в бинарном виде 0101

	i := 1                                    // в бинарном виде 0001, после сдвига 0010
	newNumber := setBitsInInt64(number, i, 1) // после побитового или 0111
	fmt.Println(newNumber)

	i = 2                                    // в бинарном виде 0010, после сдвига 0100, после побитовой инверсии 1011
	newNumber = setBitsInInt64(number, i, 0) // после побитового И 0001
	fmt.Println(newNumber)
}

func setBitsInInt64(number int64, i int, bit int) int64 {
	if bit == 0 {
		return clearBit(number, i)
	}

	return setBit(number, i)
}

func clearBit(number int64, i int) int64 {
	return number &^ (1 << i)
} // << сдвиг бита на i позиций, с помощью побитовой инверсии устанавливаем все биты в 1, кроме i, помощью логического И сбрасываем бит

func setBit(number int64, i int) int64 {
	return number | (1 << i) // << сдвиг бита на i позиций, с помощью побитовго ИЛИ устанавалием 1
}

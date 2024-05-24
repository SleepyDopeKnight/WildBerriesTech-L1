package main

import "fmt"

func main() {
	a, b := 5, 2
	fmt.Println(a, b)
	arithmetic(a, b)
	bitwiseXOR(a, b)
}

func arithmetic(a, b int) {
	a += b    // 5 + 2 = 7
	b = a - b // 7 - 5 = 2
	a -= b    // 7 - 2 = 5
	fmt.Println(a, b)
}

func bitwiseXOR(a, b int) {
	a ^= b    // 101(5) ^ 010(2) = 111(7)
	b = a ^ b // 111(7) ^ 010(2) = 101(5)
	a ^= b    // 111(7) ^ 101(5) = 010(2)
	fmt.Println(a, b)
}

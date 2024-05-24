package main

import (
	"fmt"
)

func main() {
	consistency := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 1, 0}

	subsets := make(map[int][]float64)

	for _, temperature := range consistency {
		key := int(temperature/10.0) * 10
		subsets[key] = append(subsets[key], temperature)
	}

	fmt.Println(subsets)
}

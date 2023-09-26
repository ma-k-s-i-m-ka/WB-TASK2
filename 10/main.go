package main

import (
	"fmt"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := make(map[int][]float64)
	var input int
	fmt.Print("Введите шаг: ")
	fmt.Scan(&input)

	for _, value := range arr {
		key := int(value/float64(input)) * input
		res[key] = append(res[key], value)
	}

	for key, value := range res {
		fmt.Printf("%d: %.1f\n", key, value)
	}
}

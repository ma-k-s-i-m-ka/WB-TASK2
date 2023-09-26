package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Print("Введите a и b: ")
	fmt.Scan(&a, &b)

	if a < 1048576 && b < 1048576 {
		fmt.Println("Значение меньше 1048576 ")
		return
	}
	fmt.Printf(" Сложение: %.2f,\n Вычитание: %.2f,\n Умножение: %.2f,\n Деление: %.2f\n ", a+b, a-b, a*b, a/b)
}

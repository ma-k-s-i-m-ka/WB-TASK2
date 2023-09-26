package main

import (
	"fmt"
	"strings"
)

func Check(input string) bool {
	input = strings.ToLower(input)
	m := make(map[rune]bool)

	for _, value := range input {
		if m[value] {
			return false
		}
		m[value] = true
	}
	return true
}

func main() {
	var a string
	fmt.Print("Введите строку: ")
	fmt.Scan(&a)

	res := Check(a)
	fmt.Printf("%s - %v\n", a, res)
}

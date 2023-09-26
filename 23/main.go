package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Слайс до удаления:", slice)
	var i int
	fmt.Print("Введите индекс: ")
	fmt.Scan(&i)

	if i >= 0 && i < len(slice) {
		// Удаляем элемент
		slice = append(slice[:i], slice[i+1:]...)
		fmt.Println("Слайс после удаления:", slice)
	} else {
		fmt.Println("Недопустимый индекс для удаления")
	}
}

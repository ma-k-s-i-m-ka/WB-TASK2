package main

import "fmt"

func main() {
	var arr = []int{10, 8, 9, 24, 56, 71, 98, 64}
	var a int
	fmt.Printf("Массив: %v\n", arr)
	fmt.Printf("Сортировка по возр. массива: %v\n", sortMax(arr))
	fmt.Print("Введите число: ")
	fmt.Scan(&a)
	switch {
	case a == arr[len(arr)/2]:
		fmt.Printf("Индекс равен: %v\n", len(arr)/2)
	case a < arr[len(arr)/2]:
		fmt.Printf("Поиск влево, Индекс равен: %v\n", sortLeft(a, arr))
	case a > arr[len(arr)/2]:
		fmt.Printf("Поиск вправо, Индекс равен: %v\n", sortRight(a, arr))
	}
}

func sortMax(arr []int) []int {
	for j := 0; j < 7; j++ {
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			} else {
				continue
			}
		}
	}
	return arr
}
func sortRight(a int, arr []int) int {
	for i := len(arr) / 2; i < len(arr); i++ {
		if a == arr[i] {
			return i
		}
	}
	fmt.Println("Значение не найдено")
	return 0
}

func sortLeft(a int, arr []int) int {
	for i := len(arr) / 2; i != 0; i-- {
		if a == arr[i] {
			return i
		}
	}
	fmt.Println("Значение не найдено")
	return 0
}

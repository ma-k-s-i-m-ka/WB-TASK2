package main

import "fmt"

func main() {
	var arr = []int{10, 8, 9, 24, 56, 71, 98, 64}
	fmt.Printf("Массив: %v\n", arr)
	max := arr[0]
	min := arr[0]
	fmt.Printf("Переворот массива: %v\n", sortMiddl(arr))
	fmt.Printf("Сортировка по убыв. массива: %v\n", sortMin(arr))
	fmt.Printf("Сортировка по возр. массива: %v\n", sortMax(arr))
	for i, _ := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}

	for i, _ := range arr {
		if arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Printf("max: %d, min: %d", max, min)

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

func sortMin(arr []int) []int {
	for j := 0; j < 7; j++ {
		for i := 1; i < len(arr); i++ {
			if arr[i-1] < arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			} else {
				continue
			}
		}
	}
	return arr
}

func sortMiddl(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}

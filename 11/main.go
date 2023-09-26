package main

import (
	"fmt"
)

func main() {
	s1 := []int{20, 45, 21, 16, 10}
	s2 := []int{13, 45, 8, 16, 6}
	fmt.Printf("%v\\n", Search(s1, s2))
}

func Search(a, b []int) []int {
	count := make(map[int]int)

	slice := make([]int, 0)

	for _, elem := range a {
		count[elem]++
	}

	for _, elem := range b {
		if num, ok := count[elem]; ok && num > 0 {
			count[elem]--
			slice = append(slice, elem)
		}
	}
	return slice
}

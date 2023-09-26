package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	slice := make([]int, len(arr))
	var wg sync.WaitGroup

	for i, num := range arr {
		wg.Add(1)
		go func(i, num int) {
			defer wg.Done()
			res := num * num
			slice[i] = res
		}(i, num)
	}
	wg.Wait()

	for _, res := range slice {
		fmt.Println(res)
	}

}

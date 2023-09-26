package main

import (
	"fmt"
	"sync"
)

func main() {
	ch2 := make(chan int)
	ch1 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer close(ch2)
		defer wg.Done()
		for _, num := range []int{1, 2, 3, 4, 5} {
			ch2 <- num
		}
	}()

	go func() {
		defer close(ch1)
		defer wg.Done()
		for num := range ch2 {
			res := num * 2
			ch1 <- res
		}
	}()

	go func() {
		defer wg.Done()
		for res := range ch1 {
			fmt.Println(res)
		}
	}()
	wg.Wait()
}

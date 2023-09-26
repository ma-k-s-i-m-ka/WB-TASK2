package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 10
	res := make(map[string]int, count)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			key := fmt.Sprintf("Key%d", i)
			res[key] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	for key, value := range res {
		fmt.Printf("%s: %d\n", key, value)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	resChn := make(chan int)
	doneChn := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case value, ok := <-resChn:
				if !ok {
					fmt.Println("Done")
					return
				}
				fmt.Printf("Received: %d\n", value)
			}
		}
	}()

	go func() {
		for i := 1; ; i++ {
			time.Sleep(1 * time.Second)
			select {
			case resChn <- i:
			case <-doneChn:
				close(resChn)
				return
			}
		}
	}()
	time.Sleep(10 * time.Second)
	close(doneChn)
	wg.Wait()
}

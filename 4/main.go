package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var worker = 3

	resChn := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= worker; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case value, ok := <-resChn:
					if !ok {
						fmt.Printf("Worker %d done\n", id)
						return
					}
					fmt.Printf("Worker %d received: %d\n", id, value)
				}
			}
		}(i)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for i := 1; ; i++ {
			time.Sleep(1 * time.Second)
			select {
			case resChn <- i:
			case <-sigCh:
				close(resChn)
				return
			}
		}
	}()

	wg.Wait()
}

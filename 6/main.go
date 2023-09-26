package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func gorutina_1(wg *sync.WaitGroup, exit chan struct{}) {
	defer wg.Done()
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case <-exit:
			fmt.Println("ГОРУТИНА 1  УМЕРЛА")
			return
		default:
			fmt.Println("ГОРУТИНА 1  ЖИВА")
		}
	}
}

func gorutina_2(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println("ГОРУТИНА 2  УМЕРЛА")
			return
		default:
			fmt.Println("ГОРУТИНА 2  ЖИВА")
		}
	}
}

func main() {
	exit := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go gorutina_1(&wg, exit)
	go gorutina_2(&wg, ctx)
	time.Sleep(5 * time.Second)
	exit <- struct{}{}
	cancel()
	wg.Wait()
}

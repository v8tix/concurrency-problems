package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			mu.Lock()
			time.Sleep(1 * time.Millisecond)
			mu.Unlock()
			fmt.Println("Goroutine 1 trying to acquire lock")
		}
	}()

	go func() {
		defer wg.Done()
		for {
			mu.Lock()
			time.Sleep(1 * time.Millisecond)
			mu.Unlock()
			fmt.Println("Goroutine 2 trying to acquire lock")
		}
	}()

	wg.Wait()
}

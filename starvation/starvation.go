package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		for {
			mu.Lock()
			fmt.Println("Goroutine 1 executing")
			mu.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			mu.Lock()
			fmt.Println("Goroutine 2 executing")
			mu.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			fmt.Println("Goroutine 3 waiting to execute")
			mu.Lock()
			fmt.Println("Goroutine 3 executing")
			mu.Unlock()
		}
	}()

	wg.Wait()
}

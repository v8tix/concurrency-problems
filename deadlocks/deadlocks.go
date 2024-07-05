package main

import (
	"sync"
)

func main() {
	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		defer mu1.Unlock()
		mu2.Lock()
		defer mu2.Unlock()
	}()

	go func() {
		mu2.Lock()
		defer mu2.Unlock()
		mu1.Lock()
		defer mu1.Unlock()
	}()

	// Simulate work
	select {}
}

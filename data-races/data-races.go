package main

import (
	"concurrency-problems/utils"
	"os"
	"sync"
)

var x int

func main() {
	var wg sync.WaitGroup
	logger := utils.LoadLogger()

	wg.Add(2)
	go func() {
		defer wg.Done()
		utils.RandomSleep()
		x = 1
	}()

	go func() {
		defer wg.Done()
		utils.RandomSleep()
		if x != 1 {
			logger.Error("data race", "expected", 1, "actual", x)
			os.Exit(1)
		}
	}()
	wg.Wait()
}

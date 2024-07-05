package main

import (
	"concurrency-problems/utils"
	"log/slog"
	"os"
)

var balance = 250

func main() {
	logger := utils.LoadLogger()
	go incrementBalance()
	go decrementBalance()
	go monitorBalance(logger)
	select {}
}

func incrementBalance() {
	for {
		// Sleep to make race condition more apparent
		utils.RandomSleep()
		balance = balance + 10
	}
}

func decrementBalance() {
	for {
		// Sleep to make race condition more apparent
		utils.RandomSleep()
		balance = balance - 10
	}
}

func monitorBalance(logger *slog.Logger) {
	for {
		// Sleep to make race condition more apparent
		utils.RandomSleep()
		if balance != 250 {
			logger.Error("race condition", "expected", 100, "actual", balance)
			os.Exit(1)
		}
	}
}

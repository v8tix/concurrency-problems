package utils

import (
	"log/slog"
	"math/rand"
	"os"
	"time"
)

func RandomSleep() {
	//Random sleep between 50 and 100 ms
	time.Sleep(time.Duration(rand.Intn(51)+50) * time.Millisecond)
}

func LoadLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	return logger
}

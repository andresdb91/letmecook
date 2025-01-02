package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	loggerInstance *slog.Logger
	once           sync.Once
)

func GetLogger() *slog.Logger {
	once.Do(func() {
		loggerInstance = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	})
	return loggerInstance
}

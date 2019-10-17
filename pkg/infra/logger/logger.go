package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	defaultLog "log"
	"os"
	"strings"
	"sync"
)

var (
	once   sync.Once
	logger *zap.SugaredLogger
)

func Get() *zap.SugaredLogger {
	once.Do(func() {
		log, err := newLogger()
		if err != nil {
			defaultLog.Fatalf("Can't initialize logger: %v", err)
		}
		logger = log.Sugar()
		logger.Debug("Created new logger")
	})

	return logger
}

func newLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	if lvl, exists := os.LookupEnv("LOG_LEVEL"); exists {
		lvl = strings.ToLower(lvl)
		if lvl == "debug" {
			config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		}
	}
	log, err := config.Build()
	return log, err
}

package utils

import (
	"github.com/kynmh69/study-passkey/logger"
	"go.uber.org/zap"
	"os"
)

// LookupEnv is a function to get the value of the environment variable.
func LookupEnv(key string) string {
	var value string
	if v, ok := os.LookupEnv(key); !ok {
		logger.Logger.Fatal("environment variable not found", zap.String("key", key))
	} else {
		value = v
	}
	return value
}

package utils

import (
	"fmt"
	"github.com/kynmh69/study-passkey/logger"
	"github.com/valkey-io/valkey-go"
	"go.uber.org/zap"
)

var (
	ValkeyClient valkey.Client
)

// InitValkeyClient is a function to initialize the Valkey client.
func InitValkeyClient() {
	var err error
	ipAddr := LookupEnv("VALKEY_HOST")
	port := LookupEnv("VALKEY_PORT")
	initAdds := []string{fmt.Sprintf("%s:%s", ipAddr, port)}
	ValkeyClient, err = valkey.NewClient(valkey.ClientOption{InitAddress: initAdds})
	if err != nil {
		logger.Logger.Panic(err.Error(), zap.String("host", ipAddr), zap.String("port", port))
	}
	logger.Logger.Info("Valkey client initialized", zap.String("host", ipAddr), zap.String("port", port))
}

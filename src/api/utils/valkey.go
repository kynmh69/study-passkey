package utils

import (
	"fmt"
	"github.com/kynmh69/study-passkey/consts"
	"github.com/kynmh69/study-passkey/logger"
	"github.com/valkey-io/valkey-go"
	"go.uber.org/zap"
)

// InitValkeyClient is a function to initialize the Valkey client.
func InitValkeyClient() valkey.Client {
	var err error
	ipAddr := LookupEnv(consts.ValKeyHost)
	port := LookupEnv(consts.ValKeyPort)
	initAdds := []string{fmt.Sprintf("%s:%s", ipAddr, port)}
	valkeyClient, err := valkey.NewClient(valkey.ClientOption{InitAddress: initAdds})
	if err != nil {
		logger.Logger.Panic(err.Error(), zap.String("host", ipAddr), zap.String("port", port))
	}
	logger.Logger.Info("Valkey client initialized", zap.String("host", ipAddr), zap.String("port", port))
	return valkeyClient
}

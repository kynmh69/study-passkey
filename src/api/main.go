package main

import (
	"github.com/kynmh69/study-passkey/logger"
	"github.com/kynmh69/study-passkey/server"
	"github.com/kynmh69/study-passkey/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	utils.NewSessionManager()
	utils.InitWebAuthn("StudyPass", "localhost", []string{"http://localhost:3000"})
	defer utils.Sessions.Valkey.Close()
	defer func() {
		err := utils.Client.Disconnect()
		if err != nil {
			logger.Logger.Panic("Failed to disconnect from the database", zap.Error(err))
		}
	}()
	e := echo.New()
	server.Start(e)
}

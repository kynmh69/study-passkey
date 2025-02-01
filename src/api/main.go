package main

import (
	"github.com/kynmh69/study-passkey/server"
	"github.com/kynmh69/study-passkey/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	utils.InitValkeyClient()
	defer utils.ValkeyClient.Close()
	e := echo.New()
	server.Start(e)
}

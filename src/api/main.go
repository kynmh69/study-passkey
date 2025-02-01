package main

import (
	"github.com/kynmh69/study-passkey/middleware"
	"github.com/kynmh69/study-passkey/route"
	"github.com/kynmh69/study-passkey/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	middleware.SetTimeout(e)
	middleware.SetRequestLoggerConfig(e)
	middleware.SetRecover(e)
	route.SetHandlers(e)
	server.Start(e)
}

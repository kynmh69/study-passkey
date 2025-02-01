package server

import (
	"github.com/kynmh69/study-passkey/middleware"
	"github.com/kynmh69/study-passkey/route"
	"github.com/labstack/echo/v4"
)

func Start(e *echo.Echo) {
	middleware.SetTimeout(e)
	middleware.SetRequestLoggerConfig(e)
	middleware.SetRecover(e)
	middleware.SetCors(e)
	middleware.SetSessionMiddleware(e)
	route.SetHandlers(e)
	e.Logger.Fatal(e.Start(":8080"))
}

package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetCors(e *echo.Echo) {
	e.Use(middleware.CORS())
}

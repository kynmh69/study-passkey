package route

import (
	"github.com/kynmh69/study-passkey/handler"
	"github.com/labstack/echo/v4"
)

func SetHandlers(e *echo.Echo) {
	e.GET("/users/:id", handler.GetUserById())
	e.POST("/users", handler.CreateUser())
}

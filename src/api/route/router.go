package route

import (
	"github.com/kynmh69/study-passkey/handler"
	"github.com/kynmh69/study-passkey/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetHandlers(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	api := e.Group("/api")
	v1 := api.Group("/v1")
	v1.POST("/users", handler.CreateUser())

	protectV1 := v1.Group("")
	protectV1.Use(middleware.SessionMiddleware)
	protectV1.GET("/users/:id", handler.GetUserById())
}

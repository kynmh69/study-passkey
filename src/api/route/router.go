package route

import (
	"github.com/kynmh69/study-passkey/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetHandlers(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	api := e.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/users/:id", handler.GetUserById())
	v1.POST("/users", handler.CreateUser())
}

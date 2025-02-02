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
	v1.POST("/registration/begin", handler.BeginRegistration())
	v1.POST("/registration/complete", handler.CompleteRegistration())

	protectV1 := v1.Group("")
	protectV1.Use(middleware.SessionMiddleware)
	protectV1.GET("/users/profile", handler.GetUserById())
	protectV1.POST("/logout", handler.Logout())
}

package middleware

import (
	"github.com/kynmh69/study-passkey/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// SessionMiddleware is a function to check the session.
func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionId := c.Request().Header.Get("Authorization")
		if sessionId == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Session ID is required")
		}
		sessionData, err := utils.Sessions.GetSession(sessionId)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid session ID")
		}
		c.Set("sessionId", sessionData["userId"])
		c.Set("sessionData", sessionData["metadata"])
		return next(c)
	}
}

// SetSessionMiddleware is a function to set the session middleware.
func SetSessionMiddleware(e *echo.Echo) {
	e.Use(SessionMiddleware)
}

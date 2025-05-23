package middleware

import (
	"github.com/kynmh69/study-passkey/logger"
	"github.com/kynmh69/study-passkey/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

// SessionMiddleware is a function to check the session.
func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionId := c.Request().Header.Get("Authorization")
		if sessionId == "" {
			logger.Logger.Error("Session ID is required", zap.String("sessionId", sessionId))
			return echo.NewHTTPError(http.StatusUnauthorized, "Session ID is required")
		}
		sessionData, err := utils.Sessions.GetSession(sessionId)
		if err != nil {
			logger.Logger.Error("Invalid session ID", zap.Error(err))
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid session ID")
		}
		c.Set("sessionId", sessionData["userId"])
		c.Set("sessionData", sessionData["metadata"])
		return next(c)
	}
}

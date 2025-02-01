package middleware

import (
	"github.com/kynmh69/study-passkey/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func SetRequestLoggerConfig(e *echo.Echo) {
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Logger.Info("Request",
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
				zap.String("method", v.Method),
				zap.String("path", v.RoutePath),
				zap.String("host", v.Host),
				zap.String("remote_ip", v.RemoteIP),
				zap.String("latency", v.Latency.String()),
				zap.String("user_agent", v.UserAgent),
			)
			return nil
		},
	}))
}

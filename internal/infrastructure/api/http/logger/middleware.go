package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func CustomMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogLatency:   true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context,
			values middleware.RequestLoggerValues) error {
			logrus.WithFields(logrus.Fields{
				"URI":       values.URI,
				"status":    values.Status,
				"latency":   values.Latency.Nanoseconds(),
				"requestId": values.RequestID,
			}).Info("request")

			return nil
		},
	})
}

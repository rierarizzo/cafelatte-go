package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/core/constants"
	"log/slog"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestMethod := c.Request.Method
		requestPath := c.Request.URL.Path
		requestID := c.MustGet(constants.RequestIDKey)

		start := time.Now()
		slog.Debug("Beginning request", "method", requestMethod, "path", requestPath, "requestID", requestID)

		defer func() {
			requestStatus := c.Writer.Status()

			timeElapsed := time.Since(start).Seconds()
			slog.Debug("Ending request", "method", requestMethod, "path", requestPath, "timeElapsed",
				fmt.Sprintf("%.7fs", timeElapsed), "status", requestStatus, "requestID", requestID)
		}()

		c.Next()
	}
}

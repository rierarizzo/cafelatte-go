package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/constants"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logrus.WithField(constants.RequestIDKey, params.RequestID())

		requestMethod := c.Request.Method
		requestPath := c.Request.URL.Path

		start := time.Now()
		log.WithFields(logrus.Fields{"method": requestMethod,
			"path": requestPath,
		}).Debug("Beginning request")

		c.Next()

		defer func() {
			requestStatus := c.Writer.Status()

			timeElapsed := time.Since(start).Seconds()
			log.WithFields(logrus.Fields{"method": requestMethod,
				"path":        requestPath,
				"timeElapsed": fmt.Sprintf("%.7fs", timeElapsed),
				"status":      requestStatus}).Debug("Ending request")
		}()
	}
}

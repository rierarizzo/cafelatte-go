package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
	"time"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logrus.WithField(misc.RequestIDKey, request.ID())

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

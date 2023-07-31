package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rierarizzo/cafelatte/internal/domain/constants"
	"github.com/rierarizzo/cafelatte/internal/params"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()

		c.Writer.Header().Set(constants.RequestIDHeader, requestID)
		params.SetRequestID(requestID)

		c.Next()
	}
}

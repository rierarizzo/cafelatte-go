package requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rierarizzo/cafelatte/pkg/constants/http"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()

		c.Writer.Header().Set(http.RequestIDHeader, requestID)
		request.SetRequestID(requestID)

		c.Next()
	}
}

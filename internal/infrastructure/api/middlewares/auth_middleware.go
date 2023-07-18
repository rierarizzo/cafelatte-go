package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"net/http"
	"strings"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenWithBearer := c.GetHeader("Authorization")
		if tokenWithBearer == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrTokenNotPresent)
			return
		}

		token, _ := strings.CutPrefix(tokenWithBearer, "Bearer ")
		if !utils.JWTTokenIsValid(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.ErrUnauthorizedUser)
			return
		}

		c.Next()
	}
}

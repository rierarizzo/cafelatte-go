package middlewares

import (
	"github.com/gin-gonic/gin"
	coreErrors "github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/error"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"strings"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenWithBearer := c.GetHeader("Authorization")
		if tokenWithBearer == "" {
			if error.Error(c, coreErrors.ErrTokenNotPresent) {
				return
			}
		}

		token, _ := strings.CutPrefix(tokenWithBearer, "Bearer ")

		claims, err := utils.VerifyJWTToken(token)
		if error.Error(c, err) {
			return
		}

		c.Set("userClaims", claims)
		c.Next()
	}
}

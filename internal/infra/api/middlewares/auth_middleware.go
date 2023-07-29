package middlewares

import (
	"github.com/gin-gonic/gin"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"log/slog"
	"strings"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenWithBearer := c.GetHeader("Authorization")
		if tokenWithBearer == "" {
			slog.Error(c.Error(domain.NewAppErrorWithType(domain.TokenValidationError)).Error())
			c.Abort()
			return
		}

		token, _ := strings.CutPrefix(tokenWithBearer, "Bearer ")

		claims, err := utils.VerifyJWTToken(token)
		if err != nil {
			slog.Error(c.Error(err).Error())
			c.Abort()
			return
		}

		c.Set("userClaims", claims)
		c.Next()
	}
}

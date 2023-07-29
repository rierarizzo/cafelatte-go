package middlewares

import (
	"github.com/gin-gonic/gin"
	core "github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"log/slog"
	"strings"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenWithBearer := c.GetHeader("Authorization")
		if tokenWithBearer == "" {
			slog.Error(c.Error(core.NewAppErrorWithType(core.TokenValidationError)).Error())
			return
		}

		token, _ := strings.CutPrefix(tokenWithBearer, "Bearer ")

		claims, err := utils.VerifyJWTToken(token)
		if err != nil {
			slog.Error(c.Error(err).Error())
			return
		}

		c.Set("userClaims", claims)
		c.Next()
	}
}

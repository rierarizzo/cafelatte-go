package middlewares

import (
	"github.com/gin-gonic/gin"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infra/security"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"strings"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenWithBearer := c.GetHeader("Authorization")
		if tokenWithBearer == "" {
			domainErr := domain.NewAppErrorWithType(domain.TokenValidationError)

			utils.AbortWithError(c, domainErr)
			return
		}

		token, _ := strings.CutPrefix(tokenWithBearer, "Bearer ")

		claims, err := security.VerifyJWTToken(token)
		if err != nil {
			utils.AbortWithError(c, err)
			return
		}

		c.Set("userClaims", claims)
		c.Next()
	}
}

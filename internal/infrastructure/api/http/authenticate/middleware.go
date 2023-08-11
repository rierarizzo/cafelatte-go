package authenticate

import (
	"github.com/gin-gonic/gin"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/security/jsonwebtoken"
	"github.com/rierarizzo/cafelatte/pkg/utils/http"
	"strings"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenWithBearer := c.GetHeader("Authorization")
		if tokenWithBearer == "" {
			domainErr := domain.NewAppErrorWithType(domain.TokenValidationError)

			http.AbortWithError(c, domainErr)
			return
		}

		token, _ := strings.CutPrefix(tokenWithBearer, "Bearer ")

		claims, err := jsonwebtoken.VerifyJWTToken(token)
		if err != nil {
			http.AbortWithError(c,
				domain.NewAppError(err, domain.TokenValidationError))
			return
		}

		c.Set("userClaims", claims)
		c.Next()
	}
}

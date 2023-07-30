package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/constants"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infra/security"
	"github.com/rierarizzo/cafelatte/internal/singleton"
	"github.com/sirupsen/logrus"
	"strings"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logrus.WithField(constants.RequestIDKey, singleton.RequestID())

		tokenWithBearer := c.GetHeader("Authorization")
		if tokenWithBearer == "" {
			domainErr := domain.NewAppErrorWithType(domain.TokenValidationError)

			log.Error(c.Error(domainErr).Error())
			c.Abort()
			return
		}

		token, _ := strings.CutPrefix(tokenWithBearer, "Bearer ")

		claims, err := security.VerifyJWTToken(token)
		if err != nil {
			log.Error(c.Error(err).Error())
			c.Abort()
			return
		}

		c.Set("userClaims", claims)
		c.Next()
	}
}

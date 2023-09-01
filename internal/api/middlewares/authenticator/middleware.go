package authenticator

import (
	"github.com/rierarizzo/cafelatte/internal/security/jsonwebtoken"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenWithBearer := c.Request().Header.Get("Authorization")
		if tokenWithBearer == "" {
			return domain.NewAppErrorWithType(domain.TokenValidationError)
		}

		token, _ := strings.CutPrefix(tokenWithBearer, "Bearer ")

		_, err := jsonwebtoken.VerifyJWTToken(token)
		if err != nil {
			return domain.NewAppError(err, domain.TokenValidationError)
		}

		return next(c)
	}
}

package security

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	constants2 "github.com/rierarizzo/cafelatte/internal/constants"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	sec "github.com/rierarizzo/cafelatte/internal/infrastructure/security/claims"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	invalidSigningMethodError = errors.New("signing method is not valid")
	invalidTokenError         = errors.New("token is not valid")
	expiredTokenError         = errors.New("token is expired")
	parseTokenError           = errors.New("error in parsing token")
)

func CreateJWTToken(user entities.User) (string, *domain.AppError) {
	log := logrus.WithField(constants2.RequestIDKey, params.RequestID())

	secret := []byte(os.Getenv(constants2.EnvSecretKey))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &sec.UserClaims{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Role:     user.RoleCode,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour),
			},
		},
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Error(err)

		return "", domain.NewAppError(parseTokenError,
			domain.TokenGenerationError)
	}

	return tokenString, nil
}

func VerifyJWTToken(tokenString string) (*sec.UserClaims, error) {
	log := logrus.WithField(constants2.RequestIDKey, params.RequestID())

	secret := []byte(os.Getenv(constants2.EnvSecretKey))

	var userClaims sec.UserClaims
	token, err := jwt.ParseWithClaims(tokenString, &userClaims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, invalidSigningMethodError
			}

			return secret, nil
		})
	if err != nil {
		log.Error(err)

		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, expiredTokenError
		}

		return nil, invalidTokenError
	}
	if !token.Valid {
		return nil, invalidTokenError
	}

	return &userClaims, nil
}

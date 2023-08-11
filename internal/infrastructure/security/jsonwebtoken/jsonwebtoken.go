package jsonwebtoken

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/user"
	"github.com/rierarizzo/cafelatte/pkg/constants/env"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
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

func CreateJWTToken(user user.User) (string, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	secret := []byte(os.Getenv(env.SecretKey))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
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

func VerifyJWTToken(tokenString string) (*UserClaims, error) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	secret := []byte(os.Getenv(env.SecretKey))

	var userClaims UserClaims
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

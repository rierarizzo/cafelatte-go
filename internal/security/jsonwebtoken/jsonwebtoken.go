package jsonwebtoken

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/sirupsen/logrus"
)

var (
	invalidSigningMethodError = errors.New("signing method is not valid")
	invalidTokenError         = errors.New("token is not valid")
	expiredTokenError         = errors.New("token is expired")
	parseTokenError           = errors.New("error in parsing token")
)

const SecretKey = "SECRET_KEY"

func CreateJWTToken(user domain.User) (string, *domain.AppError) {
	secret := []byte(os.Getenv(SecretKey))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		Id:       user.Id,
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
		logrus.Error(err)

		return "", domain.NewAppError(parseTokenError,
			domain.TokenGenerationError)
	}

	return tokenString, nil
}

func VerifyJWTToken(tokenString string) (*UserClaims, error) {
	secret := []byte(os.Getenv(SecretKey))

	var userClaims UserClaims
	token, err := jwt.ParseWithClaims(tokenString, &userClaims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, invalidSigningMethodError
			}

			return secret, nil
		})
	if err != nil {
		logrus.Error(err)

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

package utils

import (
	"github.com/rierarizzo/cafelatte/internal/core/constants"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"golang.org/x/crypto/bcrypt"
)

func CreateJWTToken(user entities.User) (*string, error) {
	secret := []byte(os.Getenv(constants.EnvSecretKey))

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, &entities.UserClaims{
			ID:      user.ID,
			Name:    user.Name,
			Surname: user.Surname,
			Email:   user.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt: &jwt.NumericDate{
					Time: time.Now(),
				},
				ExpiresAt: &jwt.NumericDate{
					Time: time.Now().Add(time.Hour),
				},
			},
		},
	)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	return &tokenString, nil
}

func VerifyJWTToken(tokenString string) (*entities.UserClaims, error) {
	secret := []byte(os.Getenv(constants.EnvSecretKey))

	var userClaims entities.UserClaims
	token, err := jwt.ParseWithClaims(
		tokenString, &userClaims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.WrapError(errors.ErrInvalidToken, "signing method is invalid")
			}

			return secret, nil
		},
	)
	if err != nil {
		return nil, errors.WrapError(errors.ErrInvalidToken, err.Error())
	}
	if !token.Valid {
		return nil, errors.WrapError(errors.ErrInvalidToken, "token is invalid")
	}

	return &userClaims, nil
}

func HashText(text string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	return string(bytes), nil
}

func CheckTextHash(hash string, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	return err == nil
}

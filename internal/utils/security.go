package utils

import (
	"github.com/rierarizzo/cafelatte/internal/core/constants"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/sirupsen/logrus"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}

func CreateJWTToken(user entities.User) (*string, error) {
	secret := []byte(os.Getenv(constants.EnvSecretKey))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
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
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		logrus.Errorf("error while getting token signed string: %v", err)
		return nil, errors.ErrUnexpected
	}

	return &tokenString, nil
}

func VerifyJWTToken(tokenString string) (*UserClaims, error) {
	secret := []byte(os.Getenv(constants.EnvSecretKey))

	var userClaims UserClaims
	token, err := jwt.ParseWithClaims(tokenString, &userClaims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrSignAlgorithmUnexpected
		}

		return secret, nil
	})
	if err != nil {
		logrus.Errorf("error while getting token: %v", err)
		return nil, errors.ErrUnexpected
	}
	if !token.Valid {
		return nil, errors.ErrInvalidToken
	}

	return &userClaims, nil
}

func JWTTokenIsValid(tokenString string) bool {
	_, err := VerifyJWTToken(tokenString)
	return err == nil
}

func HashText(text string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckTextHash(hash string, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	return err == nil
}

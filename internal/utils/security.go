package utils

import (
	"github.com/rierarizzo/cafelatte/internal/core"
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
	secret := []byte(os.Getenv("SECRET_KEY"))

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
				Time: time.Now().Add(time.Hour * 24),
			},
		},
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		logrus.Errorf("error while getting token string: %v", err)
		return nil, core.ErrUnexpected
	}

	return &tokenString, nil
}

func VerifyJWTToken(tokenString string) (*UserClaims, error) {
	secret := []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, core.ErrSignAlgorithmUnexpected
		}

		return secret, nil
	})
	if err != nil {
		logrus.Errorf("error while getting token: %v", err)
		return nil, core.ErrUnexpected
	}
	if !token.Valid {
		return nil, core.ErrInvalidToken
	}
	claims, ok := token.Claims.(UserClaims)
	if !ok {
		return nil, core.ErrParsingClaims
	}

	return &claims, nil
}

func HashText(text string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckTextHash(hash string, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	return err == nil
}

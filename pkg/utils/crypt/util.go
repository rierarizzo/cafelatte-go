package crypt

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func HashText(text string) (string, *domain.AppError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.NewAppErrorWithType(domain.HashGenerationError)
	}

	return string(bytes), nil
}

func CheckTextHash(hash string, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	return err == nil
}

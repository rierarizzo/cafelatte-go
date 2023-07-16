package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashText(text string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	log.Println(string(bytes))
	return string(bytes), err
}

func CheckTextHash(hash string, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	return err == nil
}

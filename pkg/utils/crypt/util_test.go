package crypt_test

import (
	"testing"

	"github.com/rierarizzo/cafelatte/pkg/utils/crypt"
	"golang.org/x/crypto/bcrypt"
)

func TestHashText(t *testing.T) {
	text := "alone again, naturally"
	hash, appErr := crypt.HashText(text)
	if appErr != nil {
		t.Errorf("Error hashing text: %v", appErr)
	}

	if len(hash) == 0 {
		t.Error("Empty hash generated")
	}

	// Verificar que el hash generado sea válido
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	if err != nil {
		t.Errorf("Invalid hash generated: %v", err)
	}
}

func TestCheckTextHash(t *testing.T) {
	text := "alone again, naturally"
	hash, _ := crypt.HashText(text)

	// Verificar que CheckTextHash retorne verdadero para una coincidencia de hash y texto
	result := crypt.CheckTextHash(hash, text)
	if result != true {
		t.Error("CheckTextHash returned false for matching hash and text")
	}

	// Verificar que CheckTextHash retorne falso para una no coincidencia de hash y texto
	incorrectText := "wrongpassword"
	result = crypt.CheckTextHash(hash, incorrectText)
	if result != false {
		t.Error("CheckTextHash returned true for non-matching hash and text")
	}
}

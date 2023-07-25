package dto

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}

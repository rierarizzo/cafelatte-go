package dto

import "github.com/rierarizzo/cafelatte/internal/core/entities"

type SignUpRequest struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (ur *SignUpRequest) ToUserCore() *entities.User {
	return &entities.User{
		Name:        ur.Name,
		Surname:     ur.Surname,
		PhoneNumber: ur.PhoneNumber,
		Email:       ur.Email,
		Password:    ur.Password,
	}
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

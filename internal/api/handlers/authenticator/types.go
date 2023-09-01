package authenticator

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserSignup struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	RoleCode    string `json:"role"`
}

func (dto *UserSignup) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Username, validation.Required, validation.Length(8, 15),
		),
		validation.Field(&dto.Name, validation.Required),
		validation.Field(&dto.Surname, validation.Required),
		validation.Field(&dto.Email, validation.Required, is.Email),
		validation.Field(&dto.Password, validation.Required),
		validation.Field(
			&dto.RoleCode, validation.Required, validation.In("A", "E", "C"),
		),
	)
}

type UserSignin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *UserSignin) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(&dto.Email, validation.Required, is.Email),
		validation.Field(&dto.Password, validation.Required),
	)
}

type AuthenticatedResponse struct {
	User struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	} `json:"user"`
	AccessToken string `json:"accessToken"`
}

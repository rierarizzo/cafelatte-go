package models

import "github.com/rierarizzo/cafelatte/internal/core/entities"

type UserModel struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Surname     string `db:"surname"`
	PhoneNumber string `db:"phone_number"`
	Email       string `db:"email"`
	Password    string `db:"password"`
}

func (um *UserModel) ToUserCore() *entities.User {
	return &entities.User{
		ID:          um.ID,
		Name:        um.Name,
		Surname:     um.Surname,
		PhoneNumber: um.PhoneNumber,
		Email:       um.Email,
		Password:    um.Password,
	}
}

func (um *UserModel) LoadFromUserCore(user entities.User) {
	um.ID = user.ID
	um.Name = user.Name
	um.Surname = user.Surname
	um.PhoneNumber = user.PhoneNumber
	um.Email = user.Email
	um.Password = user.Password
}

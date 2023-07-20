package models

import "github.com/rierarizzo/cafelatte/internal/core/entities"

type UserModel struct {
	ID          int    `db:"ID"`
	Username    string `db:"Surname"`
	Name        string `db:"Name"`
	Surname     string `db:"Surname"`
	PhoneNumber string `db:"PhoneNumber"`
	Email       string `db:"Email"`
	Password    string `db:"Password"`
	RoleCode    string `db:"RoleCode"`
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
	um.Username = user.Username
	um.Name = user.Name
	um.Surname = user.Surname
	um.PhoneNumber = user.PhoneNumber
	um.Email = user.Email
	um.Password = user.Password
	um.RoleCode = user.RoleCode
}

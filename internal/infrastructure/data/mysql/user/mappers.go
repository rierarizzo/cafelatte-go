package user

import (
	"github.com/rierarizzo/cafelatte/internal/domain/user"
)

func FromUserModelToUser(model Model) user.User {
	return user.User{
		ID:          model.ID,
		Username:    model.Username,
		Name:        model.Name,
		Surname:     model.Surname,
		PhoneNumber: model.PhoneNumber,
		Email:       model.Email,
		Password:    model.Password,
		RoleCode:    model.RoleCode,
	}
}

func FromUserModelSliceToUserSlice(models []Model) []user.User {
	var users = make([]user.User, 0)
	for _, v := range models {
		users = append(users, FromUserModelToUser(v))
	}

	return users
}

func FromUserToUserModel(user user.User) Model {
	return Model{
		ID:          user.ID,
		Username:    user.Username,
		Name:        user.Name,
		Surname:     user.Surname,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
		RoleCode:    user.RoleCode,
	}
}

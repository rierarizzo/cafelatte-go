package user

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromModelToUser(model Model) domain.User {
	return domain.User{
		Id:          model.Id,
		Username:    model.Username,
		Name:        model.Name,
		Surname:     model.Surname,
		PhoneNumber: model.PhoneNumber,
		Email:       model.Email,
		Password:    model.Password,
		RoleCode:    model.RoleCode,
	}
}

func fromModelsToUsers(models []Model) []domain.User {
	var users = make([]domain.User, 0)
	for _, v := range models {
		users = append(users, fromModelToUser(v))
	}

	return users
}

func fromUserToModel(user domain.User) Model {
	return Model{
		Id:          user.Id,
		Username:    user.Username,
		Name:        user.Name,
		Surname:     user.Surname,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
		RoleCode:    user.RoleCode,
	}
}

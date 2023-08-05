package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func FromUserModelToUser(model models.UserModel) entities.User {
	return entities.User{
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

func FromUserModelSliceToUserSlice(models []models.UserModel) []entities.User {
	var users = make([]entities.User, 0)
	for _, v := range models {
		users = append(users, FromUserModelToUser(v))
	}

	return users
}

func FromUserToUserModel(user entities.User) models.UserModel {
	return models.UserModel{
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

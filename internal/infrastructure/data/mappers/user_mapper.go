package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func FromUserModelToUser(model models.UserModel) *entities.User {
	return &entities.User{
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

func FromUserToUserModel(user entities.User) *models.UserModel {
	return &models.UserModel{
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

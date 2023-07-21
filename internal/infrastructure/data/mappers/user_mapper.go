package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func FromUserModelToUserCore(userModel models.UserModel) *entities.User {
	var addresses []entities.Address
	for _, v := range userModel.Addresses {
		addresses = append(addresses, *FromAddressModelToAddressCore(v))
	}

	var cards []entities.PaymentCard
	for _, v := range userModel.PaymentCards {
		cards = append(cards, *FromPaymentCardModelToPaymentCardCore(v))
	}

	return &entities.User{
		ID:           userModel.ID,
		Username:     userModel.Username,
		Name:         userModel.Name,
		Surname:      userModel.Surname,
		PhoneNumber:  userModel.PhoneNumber,
		Email:        userModel.Email,
		Password:     userModel.Password,
		RoleCode:     userModel.RoleCode,
		Addresses:    addresses,
		PaymentCards: cards,
	}
}

func FromUserCoreToUserModel(user entities.User) *models.UserModel {
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

package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func FromUserModelToUser(model models.UserModel) *entities.User {
	var addresses []entities.Address
	for _, v := range model.Addresses {
		addresses = append(addresses, *FromAddressModelToAddress(v))
	}

	var cards []entities.PaymentCard
	for _, v := range model.PaymentCards {
		cards = append(cards, *FromPaymentCardModelToPaymentCard(v))
	}

	return &entities.User{
		ID:           model.ID,
		Username:     model.Username,
		Name:         model.Name,
		Surname:      model.Surname,
		PhoneNumber:  model.PhoneNumber,
		Email:        model.Email,
		Password:     model.Password,
		RoleCode:     model.RoleCode,
		Addresses:    addresses,
		PaymentCards: cards,
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

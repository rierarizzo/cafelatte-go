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

func FromTemporaryUsersModelToUserSlice(tmpUsers []models.TemporaryUserModel) []entities.User {
	userMap := make(map[int]*entities.User)
	addressMap := make(map[int]bool)
	cardMap := make(map[int]bool)

	for _, v := range tmpUsers {
		userID := v.UserID

		if _, ok := userMap[userID]; !ok {
			user := entities.User{
				ID:          v.UserID,
				Username:    v.UserUsername,
				Name:        v.UserName,
				Surname:     v.UserSurname,
				PhoneNumber: v.UserPhoneNumber,
				Email:       v.UserEmail,
				Password:    v.UserPassword,
				RoleCode:    v.UserRoleCode,
			}
			userMap[userID] = &user
		}

		addressID := v.AddressID
		if !addressMap[addressID] {
			address := entities.Address{
				ID:         v.AddressID,
				Type:       v.AddressType,
				ProvinceID: v.AddressProvinceID,
				CityID:     v.AddressCityID,
				PostalCode: v.AddressPostalCode,
				Detail:     v.AddressDetail,
			}
			userMap[userID].Addresses = append(userMap[userID].Addresses, address)
			addressMap[addressID] = true
		}

		cardID := v.CardID
		if !cardMap[cardID] {
			card := entities.PaymentCard{
				ID:              v.CardID,
				Type:            v.CardType,
				Company:         v.CardCompany,
				HolderName:      v.CardHolderName,
				Number:          v.CardNumber,
				ExpirationYear:  v.CardExpirationYear,
				ExpirationMonth: v.CardExpirationMonth,
				CVV:             v.CardCVV,
			}
			userMap[userID].PaymentCards = append(userMap[userID].PaymentCards, card)
			cardMap[cardID] = true
		}
	}

	var users []entities.User
	for _, user := range userMap {
		users = append(users, *user)
	}

	return users
}

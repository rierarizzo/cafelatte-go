package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/infra/data/models"
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
		if !addressMap[int(addressID.Int64)] && addressID.Valid {
			address := entities.Address{
				ID:         int(v.AddressID.Int64),
				Type:       v.AddressType.String,
				ProvinceID: int(v.AddressProvinceID.Int64),
				CityID:     int(v.AddressCityID.Int64),
				PostalCode: v.AddressPostalCode.String,
				Detail:     v.AddressDetail.String,
			}
			userMap[userID].Addresses = append(userMap[userID].Addresses,
				address)
			addressMap[int(addressID.Int64)] = true
		}

		cardID := v.CardID
		if !cardMap[int(cardID.Int64)] && cardID.Valid {
			card := entities.PaymentCard{
				ID:              int(v.CardID.Int64),
				Type:            v.CardType.String,
				Company:         int(v.CardCompany.Int64),
				HolderName:      v.CardHolderName.String,
				Number:          v.CardNumber.String,
				ExpirationYear:  int(v.CardExpirationYear.Int64),
				ExpirationMonth: int(v.CardExpirationMonth.Int64),
				CVV:             v.CardCVV.String,
			}
			userMap[userID].PaymentCards = append(userMap[userID].PaymentCards,
				card)
			cardMap[int(cardID.Int64)] = true
		}
	}

	var users []entities.User
	for _, user := range userMap {
		users = append(users, *user)
	}

	return users
}

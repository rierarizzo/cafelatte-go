package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
	"strings"
)

func FromSignUpRequestToUserCore(signUpRequest dto.SignUpRequest) *entities.User {
	return &entities.User{
		Username:    signUpRequest.Username,
		Name:        signUpRequest.Name,
		Surname:     signUpRequest.Surname,
		PhoneNumber: signUpRequest.PhoneNumber,
		Email:       signUpRequest.Email,
		Password:    signUpRequest.Password,
		RoleCode:    signUpRequest.RoleCode,
	}
}

func FromUserCoreToUserResponse(user entities.User) *dto.UserResponse {
	var addressesResponse []dto.AddressResponse
	for _, v := range user.Addresses {
		addressesResponse = append(addressesResponse, *FromAddressCoreToAddressResponse(v))
	}

	var cardsResponse []dto.PaymentCardResponse
	for _, v := range user.PaymentCards {
		cardsResponse = append(cardsResponse, *FromPaymentCardCoreToPaymentCardResponse(v))
	}

	return &dto.UserResponse{
		ID:           user.ID,
		CompleteName: strings.Join([]string{user.Name, user.Surname}, " "),
		Username:     user.Username,
		Email:        user.Email,
		Role:         user.RoleCode,
		Addresses:    addressesResponse,
		PaymentCards: cardsResponse,
	}
}

func FromAddressCoreToAddressResponse(address entities.Address) *dto.AddressResponse {
	return &dto.AddressResponse{
		Type:       address.Type,
		ProvinceID: address.ProvinceID,
		CityID:     address.CityID,
		Detail:     address.Detail,
	}
}

func FromPaymentCardCoreToPaymentCardResponse(card entities.PaymentCard) *dto.PaymentCardResponse {
	return &dto.PaymentCardResponse{
		Type:       card.Type,
		Company:    card.Company,
		HolderName: card.HolderName,
	}
}

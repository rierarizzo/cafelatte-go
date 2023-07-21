package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
	"strings"
)

func FromSignUpReqToUser(req dto.SignUpRequest) *entities.User {
	return &entities.User{
		Username:    req.Username,
		Name:        req.Name,
		Surname:     req.Surname,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
		RoleCode:    req.RoleCode,
	}
}

func FromUserToUserRes(user entities.User) *dto.UserResponse {
	addressesRes := make([]dto.AddressResponse, 0)
	for _, v := range user.Addresses {
		addressesRes = append(addressesRes, *FromAddressToAddressRes(v))
	}

	cardsRes := make([]dto.PaymentCardResponse, 0)
	for _, v := range user.PaymentCards {
		cardsRes = append(cardsRes, *FromPaymentCardToPaymentCardRes(v))
	}

	return &dto.UserResponse{
		ID:           user.ID,
		CompleteName: strings.Join([]string{user.Name, user.Surname}, " "),
		Username:     user.Username,
		Email:        user.Email,
		Role:         user.RoleCode,
		Addresses:    addressesRes,
		PaymentCards: cardsRes,
	}
}

func FromAuthorizedUserToAuthorizationRes(authorizedUser entities.AuthorizedUser) *dto.AuthorizedUserResponse {
	userRes := FromUserToUserRes(authorizedUser.User)

	return &dto.AuthorizedUserResponse{
		User:        *userRes,
		AccessToken: authorizedUser.AccessToken,
	}
}

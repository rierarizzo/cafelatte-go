package dto

import "github.com/rierarizzo/cafelatte/internal/core/entities"

type SignUpRequest struct {
	Username     string               `json:"username"`
	Name         string               `json:"name"`
	Surname      string               `json:"surname"`
	PhoneNumber  string               `json:"phone"`
	Email        string               `json:"email"`
	Password     string               `json:"password"`
	RoleCode     string               `json:"role"`
	Addresses    []AddressRequest     `json:"addresses"`
	PaymentCards []PaymentCardRequest `json:"paymentCards"`
}

func (ur *SignUpRequest) ToUserCore() *entities.User {
	var addressesCore []entities.Address
	for _, v := range ur.Addresses {
		addressesCore = append(addressesCore, *v.ToAddressCore())
	}

	var paymentCardsCore []entities.PaymentCard
	for _, v := range ur.PaymentCards {
		paymentCardsCore = append(paymentCardsCore, *v.ToPaymentCardCore())
	}

	return &entities.User{
		Username:     ur.Username,
		Name:         ur.Name,
		Surname:      ur.Surname,
		PhoneNumber:  ur.PhoneNumber,
		Email:        ur.Email,
		Password:     ur.Password,
		RoleCode:     ur.RoleCode,
		Addresses:    addressesCore,
		PaymentCards: paymentCardsCore,
	}
}

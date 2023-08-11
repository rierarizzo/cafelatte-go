package paymentcard

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
)

type IPaymentCardService interface {
	GetCardsByUserID(userID int) ([]PaymentCard, *domain.AppError)
	AddUserPaymentCard(userID int,
		cards []PaymentCard) ([]PaymentCard, *domain.AppError)
}

type IPaymentCardRepository interface {
	SelectCardsByUserID(userID int) ([]PaymentCard, *domain.AppError)
	InsertUserPaymentCards(userID int,
		cards []PaymentCard) ([]PaymentCard, *domain.AppError)
}

package ports

import "github.com/rierarizzo/cafelatte/internal/core/entities"

type IPaymentCardService interface {
	GetCardsByUserID(userID int) ([]entities.PaymentCard, error)
	AddUserPaymentCard(
		userID int,
		cards []entities.PaymentCard,
	) ([]entities.PaymentCard, error)
}

type IPaymentCardRepository interface {
	SelectCardsByUserID(userID int) ([]entities.PaymentCard, error)
	InsertUserPaymentCards(
		userID int,
		cards []entities.PaymentCard,
	) ([]entities.PaymentCard, error)
}

package ports

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
)

type IPaymentCardService interface {
	GetCardsByUserID(userID int) ([]entities.PaymentCard, *domain.AppError)
	AddUserPaymentCard(userID int,
		cards []entities.PaymentCard) ([]entities.PaymentCard, *domain.AppError)
}

type IPaymentCardRepository interface {
	SelectCardsByUserID(userID int) ([]entities.PaymentCard, *domain.AppError)
	InsertUserPaymentCards(userID int,
		cards []entities.PaymentCard) ([]entities.PaymentCard, *domain.AppError)
}

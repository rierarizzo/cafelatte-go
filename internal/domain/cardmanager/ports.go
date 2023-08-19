package cardmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Manager interface {
	GetCardsByUserID(userID int) ([]domain.PaymentCard, *domain.AppError)
	AddUserPaymentCard(userID int,
		cards []domain.PaymentCard) ([]domain.PaymentCard, *domain.AppError)
}

type CardRepository interface {
	SelectCardsByUserID(userID int) ([]domain.PaymentCard, *domain.AppError)
	InsertUserPaymentCards(userID int,
		cards []domain.PaymentCard) ([]domain.PaymentCard, *domain.AppError)
}

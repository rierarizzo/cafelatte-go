package cardmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Manager interface {
	GetCardsByUserId(userId int) ([]domain.PaymentCard, *domain.AppError)
	AddUserCard(
		userId int,
		card domain.PaymentCard,
	) (*domain.PaymentCard, *domain.AppError)
}

type CardRepository interface {
	SelectCardsByUserID(userID int) ([]domain.PaymentCard, *domain.AppError)
	InsertUserCard(
		userId int,
		card domain.PaymentCard,
	) (*domain.PaymentCard, *domain.AppError)
}

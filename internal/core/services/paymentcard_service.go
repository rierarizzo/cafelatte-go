package services

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
)

type PaymentCardService struct {
	paymentCardRepo ports.IPaymentCardRepository
}

func (p PaymentCardService) GetCardsByUserID(userID int) ([]entities.PaymentCard, error) {
	return p.paymentCardRepo.SelectCardsByUserID(userID)
}

func (p PaymentCardService) AddUserPaymentCard(userID int, cards []entities.PaymentCard) ([]entities.PaymentCard, error) {
	for _, v := range cards {
		if err := v.ValidateExpirationDate(); err != nil {
			return nil, err
		}

		if err := v.ValidatePaymentCard(); err != nil {
			return nil, err
		}
	}

	return p.paymentCardRepo.InsertUserPaymentCards(userID, cards)
}

func NewPaymentCardService(paymentCardRepo ports.IPaymentCardRepository) *PaymentCardService {
	return &PaymentCardService{paymentCardRepo}
}

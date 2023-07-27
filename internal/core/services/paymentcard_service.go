package services

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
)

type PaymentCardService struct {
	paymentCardRepo ports.IPaymentCardRepository
}

func (s PaymentCardService) GetCardsByUserID(userID int) ([]entities.PaymentCard, error) {
	return s.paymentCardRepo.SelectCardsByUserID(userID)
}

func (s PaymentCardService) AddUserPaymentCard(userID int, cards []entities.PaymentCard) ([]entities.PaymentCard, error) {
	for _, v := range cards {
		if err := v.ValidateExpirationDate(); err != nil {
			return nil, err
		}

		if err := v.ValidatePaymentCard(); err != nil {
			return nil, err
		}
	}

	return s.paymentCardRepo.InsertUserPaymentCards(userID, cards)
}

func NewPaymentCardService(paymentCardRepo ports.IPaymentCardRepository) *PaymentCardService {
	return &PaymentCardService{paymentCardRepo}
}

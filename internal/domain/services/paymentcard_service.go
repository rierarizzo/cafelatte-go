package services

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/domain/validators"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type PaymentCardService struct {
	paymentCardRepo ports.IPaymentCardRepository
}

func (s PaymentCardService) GetCardsByUserID(userID int) ([]entities.PaymentCard, *domain.AppError) {
	cards, appErr := s.paymentCardRepo.SelectCardsByUserID(userID)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return cards, nil
}

func (s PaymentCardService) AddUserPaymentCard(userID int,
	cards []entities.PaymentCard) ([]entities.PaymentCard, *domain.AppError) {
	for k, v := range cards {
		if appErr := validators.ValidatePaymentCard(&v); appErr != nil {
			return nil, appErr
		}

		if appErr := validators.ValidateExpirationDate(&v); appErr != nil {
			return nil, appErr
		}

		hash, appErr := utils.HashText(v.Number)
		if appErr != nil {
			return nil, appErr
		}
		cards[k].Number = hash

		hash, appErr = utils.HashText(v.CVV)
		if appErr != nil {
			return nil, appErr
		}
		cards[k].CVV = hash
	}

	cards, appErr := s.paymentCardRepo.InsertUserPaymentCards(userID, cards)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return cards, nil
}

func NewPaymentCardService(paymentCardRepo ports.IPaymentCardRepository) *PaymentCardService {
	return &PaymentCardService{paymentCardRepo}
}

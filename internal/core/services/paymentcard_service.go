package services

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	core "github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type PaymentCardService struct {
	paymentCardRepo ports.IPaymentCardRepository
}

func (s PaymentCardService) GetCardsByUserID(userID int) (
	[]entities.PaymentCard,
	error,
) {
	cards, err := s.paymentCardRepo.SelectCardsByUserID(userID)
	if err != nil {
		var coreErr *core.AppError
		wrapped := errors.As(err, &coreErr)
		if (wrapped && coreErr.Type != core.NotFoundError) || !wrapped {
			return nil, core.NewAppError(err, core.UnexpectedError)
		}

		return nil, err
	}

	return cards, nil
}

func (s PaymentCardService) AddUserPaymentCard(
	userID int,
	cards []entities.PaymentCard,
) ([]entities.PaymentCard, error) {
	for k, v := range cards {
		if err := v.ValidateExpirationDate(); err != nil {
			return nil, core.NewAppError(err, core.ValidationError)
		}

		if err := v.ValidatePaymentCard(); err != nil {
			return nil, core.NewAppError(err, core.ValidationError)
		}

		hash, err := utils.HashText(v.Number)
		if err != nil {
			return nil, core.NewAppErrorWithType(core.HashGenerationError)
		}
		cards[k].Number = hash

		hash, err = utils.HashText(v.CVV)
		if err != nil {
			return nil, core.NewAppErrorWithType(core.HashGenerationError)
		}
		cards[k].CVV = hash
	}

	cards, err := s.paymentCardRepo.InsertUserPaymentCards(userID, cards)
	if err != nil {
		return nil, core.NewAppError(err, core.UnexpectedError)
	}

	return cards, nil
}

func NewPaymentCardService(paymentCardRepo ports.IPaymentCardRepository) *PaymentCardService {
	return &PaymentCardService{paymentCardRepo}
}

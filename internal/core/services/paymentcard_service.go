package services

import (
	"fmt"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	coreErrors "github.com/rierarizzo/cafelatte/internal/core/errors"
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
	return s.paymentCardRepo.SelectCardsByUserID(userID)
}

func (s PaymentCardService) AddUserPaymentCard(
	userID int,
	cards []entities.PaymentCard,
) ([]entities.PaymentCard, error) {
	for k, v := range cards {
		if err := v.ValidateExpirationDate(); err != nil {
			return nil, coreErrors.WrapError(
				err,
				fmt.Sprintf(
					"payment card with holder name '%s' is expired",
					v.HolderName,
				),
			)
		}

		if err := v.ValidatePaymentCard(); err != nil {
			return nil, coreErrors.WrapError(
				err,
				fmt.Sprintf(
					"payment card with holder name '%s' is invalid",
					v.HolderName,
				),
			)
		}

		hash, err := utils.HashText(v.Number)
		if err != nil {
			return nil, coreErrors.WrapError(err, "failed to hash card number")
		}
		cards[k].Number = hash

		hash, err = utils.HashText(v.CVV)
		if err != nil {
			return nil, coreErrors.WrapError(err, "failed to hash card CVV")
		}
		cards[k].CVV = hash
	}

	return s.paymentCardRepo.InsertUserPaymentCards(userID, cards)
}

func NewPaymentCardService(paymentCardRepo ports.IPaymentCardRepository) *PaymentCardService {
	return &PaymentCardService{paymentCardRepo}
}

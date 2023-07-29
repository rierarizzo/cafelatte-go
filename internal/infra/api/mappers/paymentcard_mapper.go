package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/infra/api/dto"
)

func FromPaymentCardReqToPaymentCard(req dto.PaymentCardRequest) *entities.PaymentCard {
	return &entities.PaymentCard{
		Type:            req.Type,
		Company:         req.Company,
		HolderName:      req.HolderName,
		Number:          req.Number,
		ExpirationYear:  req.ExpirationYear,
		ExpirationMonth: req.ExpirationMonth,
		CVV:             req.CVV,
	}
}

func FromPaymentCardToPaymentCardRes(card entities.PaymentCard) *dto.PaymentCardResponse {
	return &dto.PaymentCardResponse{
		Type:       card.Type,
		Company:    card.Company,
		HolderName: card.HolderName,
	}
}

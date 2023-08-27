package cardmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromRequestToCard(req RegisterCardRequest) domain.PaymentCard {
	return domain.PaymentCard{
		Type:            req.Type,
		Company:         req.Company,
		HolderName:      req.HolderName,
		Number:          req.Number,
		ExpirationYear:  req.ExpirationYear,
		ExpirationMonth: req.ExpirationMonth,
		CVV:             req.CVV,
	}
}

func fromCardToResponse(card *domain.PaymentCard) Response {
	return Response{
		Type:       card.Type,
		Company:    card.Company,
		HolderName: card.HolderName,
	}
}

func fromCardsToResponse(cards []domain.PaymentCard) []Response {
	res := make([]Response, 0)
	for _, v := range cards {
		res = append(res, fromCardToResponse(&v))
	}

	return res
}

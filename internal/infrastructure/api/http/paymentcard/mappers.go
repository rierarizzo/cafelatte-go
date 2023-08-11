package paymentcard

import (
	"github.com/rierarizzo/cafelatte/internal/domain/paymentcard"
)

func fromCreateRequestToCard(req CreateRequest) paymentcard.PaymentCard {
	return paymentcard.PaymentCard{
		Type:            req.Type,
		Company:         req.Company,
		HolderName:      req.HolderName,
		Number:          req.Number,
		ExpirationYear:  req.ExpirationYear,
		ExpirationMonth: req.ExpirationMonth,
		CVV:             req.CVV,
	}
}

func fromCreateRequestToCards(req []CreateRequest) []paymentcard.PaymentCard {
	cards := make([]paymentcard.PaymentCard, 0)
	for _, v := range req {
		cards = append(cards, fromCreateRequestToCard(v))
	}

	return cards
}

func fromCardToResponse(card paymentcard.PaymentCard) Response {
	return Response{
		Type:       card.Type,
		Company:    card.Company,
		HolderName: card.HolderName,
	}
}

func fromCardsToResponse(cards []paymentcard.PaymentCard) []Response {
	res := make([]Response, 0)
	for _, v := range cards {
		res = append(res, fromCardToResponse(v))
	}

	return res
}

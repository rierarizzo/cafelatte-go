package paymentcard

import (
	"github.com/rierarizzo/cafelatte/internal/domain/paymentcard"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/user"
)

func FromCardReqToCard(req user.PaymentCardRequest) paymentcard.PaymentCard {
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

func FromCardReqSliceToCardSlice(req []user.PaymentCardRequest) []paymentcard.PaymentCard {
	cards := make([]paymentcard.PaymentCard, 0)
	for _, v := range req {
		cards = append(cards, FromCardReqToCard(v))
	}

	return cards
}

func FromCardToCardRes(card paymentcard.PaymentCard) Response {
	return Response{
		Type:       card.Type,
		Company:    card.Company,
		HolderName: card.HolderName,
	}
}

func FromCardSliceToCardResSlice(cards []paymentcard.PaymentCard) []Response {
	res := make([]Response, 0)
	for _, v := range cards {
		res = append(res, FromCardToCardRes(v))
	}

	return res
}

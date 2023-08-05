package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
)

func FromCardReqToCard(req dto.PaymentCardRequest) entities.PaymentCard {
	return entities.PaymentCard{
		Type:            req.Type,
		Company:         req.Company,
		HolderName:      req.HolderName,
		Number:          req.Number,
		ExpirationYear:  req.ExpirationYear,
		ExpirationMonth: req.ExpirationMonth,
		CVV:             req.CVV,
	}
}

func FromCardReqSliceToCardSlice(req []dto.PaymentCardRequest) []entities.PaymentCard {
	cards := make([]entities.PaymentCard, 0)
	for _, v := range req {
		cards = append(cards, FromCardReqToCard(v))
	}

	return cards
}

func FromCardToCardRes(card entities.PaymentCard) dto.PaymentCardResponse {
	return dto.PaymentCardResponse{
		Type:       card.Type,
		Company:    card.Company,
		HolderName: card.HolderName,
	}
}

func FromCardSliceToCardResSlice(cards []entities.PaymentCard) []dto.PaymentCardResponse {
	res := make([]dto.PaymentCardResponse, 0)
	for _, v := range cards {
		res = append(res, FromCardToCardRes(v))
	}

	return res
}

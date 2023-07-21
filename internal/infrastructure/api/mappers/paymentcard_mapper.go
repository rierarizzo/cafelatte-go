package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
)

func FromPaymentCardRequestSliceToPaymentCardCoreSlice(req []dto.PaymentCardRequest) []entities.PaymentCard {
	var cards []entities.PaymentCard
	for _, k := range req {
		cards = append(cards, entities.PaymentCard{
			Type:           k.Type,
			Company:        k.Company,
			HolderName:     k.HolderName,
			Number:         k.Number,
			ExpirationDate: k.ExpirationDate,
			CVV:            k.CVV,
		})
	}

	return cards
}

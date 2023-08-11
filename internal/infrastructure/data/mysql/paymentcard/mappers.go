package paymentcard

import (
	"database/sql"
	"github.com/rierarizzo/cafelatte/internal/domain/paymentcard"
)

func FromCardToCardModel(card paymentcard.PaymentCard) Model {
	return Model{
		ID:              sql.NullInt64{Int64: int64(card.ID)},
		Type:            card.Type,
		Company:         card.Company,
		HolderName:      card.HolderName,
		Number:          card.Number,
		ExpirationYear:  card.ExpirationYear,
		ExpirationMonth: card.ExpirationMonth,
		CVV:             card.CVV,
	}
}

func FromCardModelToCard(model Model) paymentcard.PaymentCard {
	return paymentcard.PaymentCard{
		ID:              int(model.ID.Int64),
		Type:            model.Type,
		Company:         model.Company,
		HolderName:      model.HolderName,
		Number:          model.Number,
		ExpirationYear:  model.ExpirationYear,
		ExpirationMonth: model.ExpirationMonth,
		CVV:             model.CVV,
	}
}

func FromCardModelSliceToCardSlice(cardsModel []Model) []paymentcard.PaymentCard {
	var cards = make([]paymentcard.PaymentCard, 0)
	for _, v := range cardsModel {
		cards = append(cards, FromCardModelToCard(v))
	}

	return cards
}

package card

import (
	"database/sql"

	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromCardToModel(card domain.PaymentCard) Model {
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

func fromModelToCard(model Model) domain.PaymentCard {
	return domain.PaymentCard{
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

func fromModelsToCards(cardsModel []Model) []domain.PaymentCard {
	var cards = make([]domain.PaymentCard, 0)
	for _, v := range cardsModel {
		cards = append(cards, fromModelToCard(v))
	}

	return cards
}

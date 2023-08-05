package mappers

import (
	"database/sql"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func FromCardToCardModel(card entities.PaymentCard) models.PaymentCardModel {
	return models.PaymentCardModel{
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

func FromCardModelToCard(model models.PaymentCardModel) entities.PaymentCard {
	return entities.PaymentCard{
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

func FromCardModelSliceToCardSlice(cardsModel []models.PaymentCardModel) []entities.PaymentCard {
	var cards = make([]entities.PaymentCard, 0)
	for _, v := range cardsModel {
		cards = append(cards, FromCardModelToCard(v))
	}

	return cards
}

package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func FromPaymentCardToPaymentCardModel(card entities.PaymentCard) *models.PaymentCardModel {
	return &models.PaymentCardModel{
		Type:            card.Type,
		Company:         card.Company,
		HolderName:      card.HolderName,
		Number:          card.Number,
		ExpirationYear:  card.ExpirationYear,
		ExpirationMonth: card.ExpirationMonth,
		CVV:             card.CVV,
	}
}

func FromPaymentCardModelToPaymentCard(model models.PaymentCardModel) *entities.PaymentCard {
	return &entities.PaymentCard{
		ID:              model.ID,
		Type:            model.Type,
		Company:         model.Company,
		HolderName:      model.HolderName,
		Number:          model.Number,
		ExpirationYear:  model.ExpirationYear,
		ExpirationMonth: model.ExpirationMonth,
		CVV:             model.CVV,
	}
}

package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func PaymentCardCoreToPaymentCardModel(card entities.PaymentCard) *models.PaymentCardModel {
	return &models.PaymentCardModel{
		Type:           card.Type,
		Company:        card.Company,
		HolderName:     card.HolderName,
		Number:         card.Number,
		ExpirationDate: card.ExpirationDate,
		CVV:            card.CVV,
	}
}

func PaymentCardModelToPaymentCardCore(cardModel models.PaymentCardModel) *entities.PaymentCard {
	return &entities.PaymentCard{
		ID:             cardModel.ID,
		Type:           cardModel.Type,
		Company:        cardModel.Company,
		HolderName:     cardModel.HolderName,
		Number:         cardModel.Number,
		ExpirationDate: cardModel.ExpirationDate,
		CVV:            cardModel.CVV,
	}
}

package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func PaymentCardCoreToPaymentCardModel(card entities.PaymentCard, userID int) *models.PaymentCardModel {
	return &models.PaymentCardModel{
		Type:           card.Type,
		UserID:         userID,
		Company:        card.Company,
		Issuer:         card.Issuer,
		HolderName:     card.HolderName,
		Number:         card.Number,
		ExpirationDate: card.ExpirationDate,
		CVV:            card.CVV,
	}
}

package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
)

func FromCardReqToCardCore(req dto.PaymentCardRequest) *entities.PaymentCard {
	return &entities.PaymentCard{
		Type:           req.Type,
		Company:        req.Company,
		HolderName:     req.HolderName,
		Number:         req.Number,
		ExpirationDate: req.ExpirationDate,
		CVV:            req.CVV,
	}
}

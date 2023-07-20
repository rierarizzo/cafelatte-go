package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func AddressCoreToAddressModel(address entities.Address, userID int) *models.AddressModel {
	return &models.AddressModel{
		Type:       address.Type,
		UserID:     userID,
		ProvinceID: address.ProvinceID,
		CityID:     address.CityID,
		PostalCode: address.PostalCode,
		Detail:     address.Detail,
	}
}

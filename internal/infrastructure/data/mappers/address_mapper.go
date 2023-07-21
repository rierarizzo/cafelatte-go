package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func FromAddressToAddressModel(address entities.Address) *models.AddressModel {
	return &models.AddressModel{
		Type:       address.Type,
		ProvinceID: address.ProvinceID,
		CityID:     address.CityID,
		PostalCode: address.PostalCode,
		Detail:     address.Detail,
	}
}

func FromAddressModelToAddress(model models.AddressModel) *entities.Address {
	return &entities.Address{
		ID:         model.ID,
		Type:       model.Type,
		ProvinceID: model.ProvinceID,
		CityID:     model.CityID,
		PostalCode: model.PostalCode,
		Detail:     model.Detail,
	}
}

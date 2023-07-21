package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func AddressCoreToAddressModel(address entities.Address) *models.AddressModel {
	return &models.AddressModel{
		Type:       address.Type,
		ProvinceID: address.ProvinceID,
		CityID:     address.CityID,
		PostalCode: address.PostalCode,
		Detail:     address.Detail,
	}
}

func AddressModelToAddressCore(addressModel models.AddressModel) *entities.Address {
	return &entities.Address{
		ID:         addressModel.ID,
		Type:       addressModel.Type,
		ProvinceID: addressModel.ProvinceID,
		CityID:     addressModel.CityID,
		PostalCode: addressModel.PostalCode,
		Detail:     addressModel.Detail,
	}
}

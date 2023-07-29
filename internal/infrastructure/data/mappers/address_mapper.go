package mappers

import (
	"database/sql"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func FromAddressToAddressModel(address entities.Address) *models.AddressModel {
	return &models.AddressModel{
		ID:         sql.NullInt64{Int64: int64(address.ID)},
		Type:       address.Type,
		ProvinceID: address.ProvinceID,
		CityID:     address.CityID,
		PostalCode: address.PostalCode,
		Detail:     address.Detail,
	}
}

func FromAddressModelToAddress(model models.AddressModel) *entities.Address {
	return &entities.Address{
		ID:         int(model.ID.Int64),
		Type:       model.Type,
		ProvinceID: model.ProvinceID,
		CityID:     model.CityID,
		PostalCode: model.PostalCode,
		Detail:     model.Detail,
	}
}

package address

import (
	"database/sql"

	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromAddressToModel(address domain.Address) Model {
	return Model{
		ID:         sql.NullInt64{Int64: int64(address.ID)},
		Type:       address.Type,
		ProvinceID: address.ProvinceID,
		CityID:     address.CityID,
		PostalCode: address.PostalCode,
		Detail:     address.Detail,
	}
}

func fromModelToAddress(model Model) domain.Address {
	return domain.Address{
		ID:         int(model.ID.Int64),
		Type:       model.Type,
		ProvinceID: model.ProvinceID,
		CityID:     model.CityID,
		PostalCode: model.PostalCode,
		Detail:     model.Detail,
	}
}

func fromModelsToAddresses(addressesModel []Model) []domain.Address {
	var addresses = make([]domain.Address, 0)
	for _, v := range addressesModel {
		addresses = append(addresses, fromModelToAddress(v))
	}

	return addresses
}

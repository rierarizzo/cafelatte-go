package address

import (
	"database/sql"
	"github.com/rierarizzo/cafelatte/internal/domain/address"
)

func ToModel(address address.Address) Model {
	return Model{
		ID:         sql.NullInt64{Int64: int64(address.ID)},
		Type:       address.Type,
		ProvinceID: address.ProvinceID,
		CityID:     address.CityID,
		PostalCode: address.PostalCode,
		Detail:     address.Detail,
	}
}

func ModelToAddress(model Model) address.Address {
	return address.Address{
		ID:         int(model.ID.Int64),
		Type:       model.Type,
		ProvinceID: model.ProvinceID,
		CityID:     model.CityID,
		PostalCode: model.PostalCode,
		Detail:     model.Detail,
	}
}

func ModelSliceToAddresses(addressesModel []Model) []address.Address {
	var addresses = make([]address.Address, 0)
	for _, v := range addressesModel {
		addresses = append(addresses, ModelToAddress(v))
	}

	return addresses
}

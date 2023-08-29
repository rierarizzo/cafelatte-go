package address

import (
	"database/sql"

	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromAddressToModel(address domain.Address) Model {
	return Model{
		Id:         sql.NullInt64{Int64: int64(address.Id)},
		Type:       address.Type,
		ProvinceId: address.ProvinceId,
		CityId:     address.CityId,
		PostalCode: address.PostalCode,
		Detail:     address.Detail,
	}
}

func fromModelToAddress(model Model) domain.Address {
	return domain.Address{
		Id:         int(model.Id.Int64),
		Type:       model.Type,
		ProvinceId: model.ProvinceId,
		CityId:     model.CityId,
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

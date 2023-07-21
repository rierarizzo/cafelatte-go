package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
)

func FromAddressRequestSliceToAddressCoreSlice(req []dto.AddressRequest) []entities.Address {
	var addresses []entities.Address
	for _, k := range req {
		addresses = append(addresses, entities.Address{
			Type:       k.Type,
			ProvinceID: k.ProvinceID,
			CityID:     k.CityID,
			PostalCode: k.PostalCode,
			Detail:     k.Detail,
		})
	}

	return addresses
}

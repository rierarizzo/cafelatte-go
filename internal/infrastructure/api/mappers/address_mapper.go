package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
)

func FromAddressReqToAddress(req dto.AddressRequest) entities.Address {
	return entities.Address{
		Type:       req.Type,
		ProvinceID: req.ProvinceID,
		CityID:     req.CityID,
		PostalCode: req.PostalCode,
		Detail:     req.Detail,
	}
}

func FromAddressReqSliceToAddressSlice(req []dto.AddressRequest) []entities.Address {
	addresses := make([]entities.Address, 0)
	for _, v := range req {
		addresses = append(addresses, FromAddressReqToAddress(v))
	}

	return addresses
}

func FromAddressToAddressRes(address entities.Address) dto.AddressResponse {
	return dto.AddressResponse{
		Type:       address.Type,
		ProvinceID: address.ProvinceID,
		CityID:     address.CityID,
		Detail:     address.Detail,
	}
}

func FromAddressSliceToAddressResSlice(addresses []entities.Address) []dto.AddressResponse {
	res := make([]dto.AddressResponse, 0)
	for _, v := range addresses {
		res = append(res, FromAddressToAddressRes(v))
	}

	return res
}

package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
)

func FromAddressReqToAddress(req dto.AddressRequest) *entities.Address {
	return &entities.Address{
		Type:       req.Type,
		ProvinceID: req.ProvinceID,
		CityID:     req.CityID,
		PostalCode: req.PostalCode,
		Detail:     req.Detail,
	}
}

func FromAddressToAddressRes(address entities.Address) *dto.AddressResponse {
	return &dto.AddressResponse{
		Type:       address.Type,
		ProvinceID: address.ProvinceID,
		CityID:     address.CityID,
		Detail:     address.Detail,
	}
}

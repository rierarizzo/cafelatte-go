package addressmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromRequestToAddress(req AddressCreate) domain.Address {
	return domain.Address{
		Type:       req.Type,
		ProvinceId: req.ProvinceId,
		CityId:     req.CityId,
		PostalCode: req.PostalCode,
		Detail:     req.Detail,
	}
}

func fromAddressToResponse(address *domain.Address) AddressResponse {
	return AddressResponse{
		Type:       address.Type,
		ProvinceId: address.ProvinceId,
		CityId:     address.CityId,
		Detail:     address.Detail,
	}
}

func fromAddressesToResponse(addresses []domain.Address) []AddressResponse {
	res := make([]AddressResponse, 0)
	for _, v := range addresses {
		res = append(res, fromAddressToResponse(&v))
	}

	return res
}

package addressmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromCreateRequestToAddress(req CreateRequest) domain.Address {
	return domain.Address{
		Type:       req.Type,
		ProvinceID: req.ProvinceID,
		CityID:     req.CityID,
		PostalCode: req.PostalCode,
		Detail:     req.Detail,
	}
}

func fromCreateRequestToAddresses(req []CreateRequest) []domain.Address {
	addresses := make([]domain.Address, 0)
	for _, v := range req {
		addresses = append(addresses, fromCreateRequestToAddress(v))
	}

	return addresses
}

func fromAddressToResponse(address domain.Address) Response {
	return Response{
		Type:       address.Type,
		ProvinceID: address.ProvinceID,
		CityID:     address.CityID,
		Detail:     address.Detail,
	}
}

func fromAddressesToResponse(addresses []domain.Address) []Response {
	res := make([]Response, 0)
	for _, v := range addresses {
		res = append(res, fromAddressToResponse(v))
	}

	return res
}

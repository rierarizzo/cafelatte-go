package address

import (
	"github.com/rierarizzo/cafelatte/internal/domain/address"
)

func fromCreateRequestToAddress(req CreateRequest) address.Address {
	return address.Address{
		Type:       req.Type,
		ProvinceID: req.ProvinceID,
		CityID:     req.CityID,
		PostalCode: req.PostalCode,
		Detail:     req.Detail,
	}
}

func fromCreateRequestToAddresses(req []CreateRequest) []address.Address {
	addresses := make([]address.Address, 0)
	for _, v := range req {
		addresses = append(addresses, fromCreateRequestToAddress(v))
	}

	return addresses
}

func fromAddressToResponse(address address.Address) Response {
	return Response{
		Type:       address.Type,
		ProvinceID: address.ProvinceID,
		CityID:     address.CityID,
		Detail:     address.Detail,
	}
}

func fromAddressesToResponse(addresses []address.Address) []Response {
	res := make([]Response, 0)
	for _, v := range addresses {
		res = append(res, fromAddressToResponse(v))
	}

	return res
}

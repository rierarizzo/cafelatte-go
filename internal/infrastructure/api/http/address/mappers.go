package address

import (
	"github.com/rierarizzo/cafelatte/internal/domain/address"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/user"
)

func FromAddressReqToAddress(req user.AddressRequest) address.Address {
	return address.Address{
		Type:       req.Type,
		ProvinceID: req.ProvinceID,
		CityID:     req.CityID,
		PostalCode: req.PostalCode,
		Detail:     req.Detail,
	}
}

func FromAddressReqSliceToAddressSlice(req []user.AddressRequest) []address.Address {
	addresses := make([]address.Address, 0)
	for _, v := range req {
		addresses = append(addresses, FromAddressReqToAddress(v))
	}

	return addresses
}

func FromAddressToAddressRes(address address.Address) Response {
	return Response{
		Type:       address.Type,
		ProvinceID: address.ProvinceID,
		CityID:     address.CityID,
		Detail:     address.Detail,
	}
}

func FromAddressSliceToAddressResSlice(addresses []address.Address) []Response {
	res := make([]Response, 0)
	for _, v := range addresses {
		res = append(res, FromAddressToAddressRes(v))
	}

	return res
}

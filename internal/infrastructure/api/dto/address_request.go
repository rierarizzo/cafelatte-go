package dto

import "github.com/rierarizzo/cafelatte/internal/core/entities"

type AddressRequest struct {
	Type       string `json:"type"`
	ProvinceID int    `json:"province"`
	CityID     int    `json:"city"`
	PostalCode string `json:"postalCode"`
	Detail     string `json:"detail"`
}

func (ar *AddressRequest) ToAddressCore() *entities.Address {
	return &entities.Address{
		Type:       ar.Type,
		ProvinceID: ar.ProvinceID,
		CityID:     ar.CityID,
		PostalCode: ar.PostalCode,
		Detail:     ar.Detail,
	}
}

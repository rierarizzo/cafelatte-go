package addressmanager

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

type AddressCreate struct {
	Type       string `json:"type"`
	ProvinceId int    `json:"provinceId"`
	CityId     int    `json:"cityId"`
	PostalCode string `json:"postalCode"`
	Detail     string `json:"detail"`
}

func (dto *AddressCreate) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Type, validation.Required),
		validation.Field(&dto.ProvinceId, validation.Required),
		validation.Field(&dto.CityId, validation.Required),
		validation.Field(&dto.PostalCode,
			validation.Length(6, 6),
			validation.Match(regexp.MustCompile("^[0-9]+$"))))
}

type AddressResponse struct {
	Type       string `json:"type"`
	ProvinceId int    `json:"provinceId"`
	CityId     int    `json:"cityId"`
	Detail     string `json:"detail"`
}

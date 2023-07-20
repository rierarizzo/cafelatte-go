package models

import "github.com/rierarizzo/cafelatte/internal/core/entities"

type AddressModel struct {
	ID         int    `db:"ID"`
	Type       string `db:"Type"`
	UserID     int    `db:"UserID"`
	ProvinceID int    `db:"ProvinceID"`
	CityID     int    `db:"CityID"`
	PostalCode string `db:"PostalCode"`
	Detail     string `db:"Detail"`
	Enabled    bool   `db:"Enabled"`
}

func (am *AddressModel) LoadFromAddressCore(address entities.Address) {
	am.ID = address.ID
	am.Type = address.Type
	am.ProvinceID = address.ProvinceID
	am.CityID = address.CityID
	am.PostalCode = address.PostalCode
	am.Detail = address.Detail
}

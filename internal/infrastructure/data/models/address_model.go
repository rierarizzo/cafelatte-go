package models

import "time"

type AddressModel struct {
	ID         int       `db:"ID"`
	Type       string    `db:"Type"`
	UserID     int       `db:"UserID"`
	ProvinceID int       `db:"ProvinceID"`
	CityID     int       `db:"CityID"`
	PostalCode string    `db:"PostalCode"`
	Detail     string    `db:"Detail"`
	Status     bool      `db:"Status"`
	CreatedAt  time.Time `db:"CreatedAt"`
	UpdatedAt  time.Time `db:"UpdatedAt"`
}

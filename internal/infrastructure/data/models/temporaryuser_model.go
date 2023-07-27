package models

import (
	"time"
)

type TemporaryUserModel struct {
	UserID          int       `db:"UserID"`
	UserUsername    string    `db:"UserUsername"`
	UserName        string    `db:"UserName"`
	UserSurname     string    `db:"UserSurname"`
	UserPhoneNumber string    `db:"UserPhoneNumber"`
	UserEmail       string    `db:"UserEmail"`
	UserPassword    string    `db:"UserPassword"`
	UserRoleCode    string    `db:"UserRoleCode"`
	UserStatus      bool      `db:"UserStatus"`
	UserCreatedAt   time.Time `db:"UserCreatedAt"`
	UserUpdatedAt   time.Time `db:"UserUpdatedAt"`

	AddressID         int       `db:"AddressID"`
	AddressType       string    `db:"AddressType"`
	AddressProvinceID int       `db:"AddressProvinceID"`
	AddressCityID     int       `db:"AddressCityID"`
	AddressPostalCode string    `db:"AddressPostalCode"`
	AddressDetail     string    `db:"AddressDetail"`
	AddressStatus     bool      `db:"AddressStatus"`
	AddressCreatedAt  time.Time `db:"AddressCreatedAt"`
	AddressUpdatedAt  time.Time `db:"AddressUpdatedAt"`

	CardID              int       `db:"CardID"`
	CardType            string    `db:"CardType"`
	CardCompany         int       `db:"CardCompany"`
	CardHolderName      string    `db:"CardHolderName"`
	CardNumber          string    `db:"CardNumber"`
	CardExpirationYear  int       `db:"CardExpirationYear"`
	CardExpirationMonth int       `db:"CardExpirationMonth"`
	CardCVV             string    `db:"CardCVV"`
	CardStatus          bool      `db:"CardStatus"`
	CardCreatedAt       time.Time `db:"CardCreatedAt"`
	CardUpdatedAt       time.Time `db:"CardUpdatedAt"`
}

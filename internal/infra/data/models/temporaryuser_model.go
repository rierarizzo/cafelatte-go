package models

import (
	"database/sql"
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

	AddressID         sql.NullInt64  `db:"AddressID"`
	AddressType       sql.NullString `db:"AddressType"`
	AddressProvinceID sql.NullInt64  `db:"AddressProvinceID"`
	AddressCityID     sql.NullInt64  `db:"AddressCityID"`
	AddressPostalCode sql.NullString `db:"AddressPostalCode"`
	AddressDetail     sql.NullString `db:"AddressDetail"`
	AddressStatus     sql.NullBool   `db:"AddressStatus"`
	AddressCreatedAt  sql.NullTime   `db:"AddressCreatedAt"`
	AddressUpdatedAt  sql.NullTime   `db:"AddressUpdatedAt"`

	CardID              sql.NullInt64  `db:"CardID"`
	CardType            sql.NullString `db:"CardType"`
	CardCompany         sql.NullInt64  `db:"CardCompany"`
	CardHolderName      sql.NullString `db:"CardHolderName"`
	CardNumber          sql.NullString `db:"CardNumber"`
	CardExpirationYear  sql.NullInt64  `db:"CardExpirationYear"`
	CardExpirationMonth sql.NullInt64  `db:"CardExpirationMonth"`
	CardCVV             sql.NullString `db:"CardCVV"`
	CardStatus          sql.NullBool   `db:"CardStatus"`
	CardCreatedAt       sql.NullTime   `db:"CardCreatedAt"`
	CardUpdatedAt       sql.NullTime   `db:"CardUpdatedAt"`
}

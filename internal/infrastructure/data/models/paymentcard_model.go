package models

import (
	"time"
)

type PaymentCardModel struct {
	ID             int       `db:"ID"`
	Type           string    `db:"Type"`
	UserID         int       `db:"UserID"`
	Company        int       `db:"Company"`
	Issuer         int       `db:"Issuer"`
	HolderName     string    `db:"HolderName"`
	Number         string    `db:"Number"`
	ExpirationDate time.Time `db:"ExpirationDate"`
	CVV            string    `db:"CVV"`
	Enabled        bool      `db:"Enabled"`
}

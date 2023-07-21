package models

import (
	"time"
)

type PaymentCardModel struct {
	ID             int       `db:"ID"`
	Type           string    `db:"Type"`
	Company        int       `db:"Company"`
	HolderName     string    `db:"HolderName"`
	Number         string    `db:"Number"`
	ExpirationDate time.Time `db:"ExpirationDate"`
	CVV            string    `db:"CVV"`
	Status         bool      `db:"Status"`
	CreatedAt      time.Time `db:"CreatedAt"`
	UpdatedAt      time.Time `db:"UpdatedAt"`
}
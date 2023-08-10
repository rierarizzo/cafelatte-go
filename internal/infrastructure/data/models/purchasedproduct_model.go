package models

import (
	"time"
)

type PurchasedProductModel struct {
	ID        int       `db:"ID"`
	OrderID   int       `db:"OrderID"`
	ProductID int       `db:"ProductID"`
	Quantity  int       `db:"Quantity"`
	CreatedAt time.Time `db:"CreatedAt"`
	UpdatedAt time.Time `db:"UpdatedAt"`
}

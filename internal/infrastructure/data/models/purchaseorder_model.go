package models

import (
	"database/sql"
	"time"
)

type PurchaseOrderModel struct {
	ID                int             `db:"ID"`
	UserID            int             `db:"UserID"`
	ShippingAddressID int             `db:"ShippingAddressID"`
	PaymentMethodID   int             `db:"PaymentMethodID"`
	Notes             sql.NullString  `db:"Notes"`
	TotalAmount       sql.NullFloat64 `db:"TotalAmount"`
	OrderedAt         time.Time       `db:"OrderedAt"`
	OrderStatus       string          `db:"OrderStatus"`
	CreatedAt         time.Time       `db:"CreatedAt"`
	UpdatedAt         time.Time       `db:"UpdatedAt"`
}

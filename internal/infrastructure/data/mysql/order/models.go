package order

import (
	"database/sql"
	"time"
)

type Model struct {
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

type PurchasedProductModel struct {
	ID        int       `db:"ID"`
	OrderID   int       `db:"OrderID"`
	ProductID int       `db:"ProductID"`
	Quantity  int       `db:"Quantity"`
	CreatedAt time.Time `db:"CreatedAt"`
	UpdatedAt time.Time `db:"UpdatedAt"`
}

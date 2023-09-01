package order

import (
	"database/sql"
	"time"
)

type Model struct {
	Id                int             `db:"Id"`
	UserId            int             `db:"UserId"`
	ShippingAddressId int             `db:"ShippingAddressId"`
	PaymentMethodId   int             `db:"PaymentMethodId"`
	Notes             sql.NullString  `db:"Notes"`
	TotalAmount       sql.NullFloat64 `db:"TotalAmount"`
	OrderedAt         time.Time       `db:"OrderedAt"`
	OrderStatus       string          `db:"OrderStatus"`
	CreatedAt         time.Time       `db:"CreatedAt"`
	UpdatedAt         time.Time       `db:"UpdatedAt"`
}

type ProductInOrderModel struct {
	Id        int       `db:"Id"`
	OrderId   int       `db:"OrderId"`
	ProductId int       `db:"ProductId"`
	Quantity  int       `db:"Quantity"`
	CreatedAt time.Time `db:"CreatedAt"`
	UpdatedAt time.Time `db:"UpdatedAt"`
}

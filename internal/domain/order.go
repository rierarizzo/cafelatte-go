package domain

import (
	"time"
)

type Order struct {
	Id                int
	UserId            int
	ShippingAddressId int
	PaymentMethodId   int
	Notes             string
	TotalAmount       float64
	OrderedAt         time.Time
	OrderStatus       string
	Products          []ProductInOrder
}

type ProductInOrder struct {
	Id        int
	OrderId   int
	ProductId int
	Quantity  int
}

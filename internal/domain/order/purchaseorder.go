package order

import (
	"time"
)

type Order struct {
	ID                int
	UserID            int
	ShippingAddressID int
	PaymentMethodID   int
	Notes             string
	TotalAmount       float64
	OrderedAt         time.Time
	OrderStatus       string
	PurchasedProducts []PurchasedProduct
}

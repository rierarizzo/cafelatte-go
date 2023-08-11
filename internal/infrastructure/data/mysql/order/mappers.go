package order

import (
	"database/sql"
	"github.com/rierarizzo/cafelatte/internal/domain/order"
)

func ToModel(order order.Order) PurchaseOrderModel {
	return PurchaseOrderModel{
		ID:                order.ID,
		UserID:            order.UserID,
		ShippingAddressID: order.ShippingAddressID,
		PaymentMethodID:   order.PaymentMethodID,
		Notes:             sql.NullString{String: order.Notes},
		TotalAmount:       sql.NullFloat64{Float64: order.TotalAmount},
		OrderedAt:         order.OrderedAt,
		OrderStatus:       order.OrderStatus,
	}
}

func PurchasedProductToModel(product order.PurchasedProduct) PurchasedProductModel {
	return PurchasedProductModel{
		ID:        product.ID,
		OrderID:   product.OrderID,
		ProductID: product.ProductID,
		Quantity:  product.Quantity,
	}
}

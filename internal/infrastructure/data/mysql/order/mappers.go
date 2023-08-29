package order

import (
	"database/sql"

	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromOrderToModel(order domain.Order) Model {
	return Model{
		Id:                order.Id,
		UserId:            order.UserId,
		ShippingAddressId: order.ShippingAddressId,
		PaymentMethodId:   order.PaymentMethodId,
		Notes:             sql.NullString{String: order.Notes},
		TotalAmount:       sql.NullFloat64{Float64: order.TotalAmount},
		OrderedAt:         order.OrderedAt,
		OrderStatus:       order.OrderStatus,
	}
}

func fromProductInOrderToModel(product domain.ProductInOrder) ProductInOrderModel {
	return ProductInOrderModel{
		Id:        product.Id,
		OrderId:   product.OrderId,
		ProductId: product.ProductId,
		Quantity:  product.Quantity,
	}
}

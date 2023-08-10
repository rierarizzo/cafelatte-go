package mappers

import (
	"database/sql"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func OrderToModel(order entities.PurchaseOrder) models.PurchaseOrderModel {
	return models.PurchaseOrderModel{
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

func PurchasedProductToModel(product entities.PurchasedProduct) models.PurchasedProductModel {
	return models.PurchasedProductModel{
		ID:        product.ID,
		OrderID:   product.OrderID,
		ProductID: product.ProductID,
		Quantity:  product.Quantity,
	}
}

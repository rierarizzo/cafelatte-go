package purchase

import (
	"github.com/rierarizzo/cafelatte/internal/domain/order"
)

func fromCreateOrderRequestToOrder(req CreateOrderRequest) order.Order {
	products := make([]order.PurchasedProduct, 0)
	for _, v := range req.PurchasedProducts {
		products = append(products, order.PurchasedProduct{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		})
	}

	return order.Order{
		UserID:            req.UserID,
		ShippingAddressID: req.AddressID,
		PaymentMethodID:   req.PaymentMethodID,
		Notes:             req.Notes,
		PurchasedProducts: products,
	}
}

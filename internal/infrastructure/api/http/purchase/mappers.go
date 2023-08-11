package purchase

import (
	"github.com/rierarizzo/cafelatte/internal/domain/order"
)

func requestToPurchasedProduct(req PurchasedProduct) order.PurchasedProduct {
	return order.PurchasedProduct{
		ID:        req.ProductID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
}

func RequestToOrder(req OrderRequest) order.Order {
	products := make([]order.PurchasedProduct, 0)
	for _, v := range req.PurchasedProducts {
		products = append(products, requestToPurchasedProduct(v))
	}

	return order.Order{
		UserID:            req.UserID,
		ShippingAddressID: req.AddressID,
		PaymentMethodID:   req.PaymentMethodID,
		Notes:             req.Notes,
		PurchasedProducts: products,
	}
}

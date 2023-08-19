package productpurchaser

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromCreateOrderRequestToOrder(req CreateOrderRequest) domain.Order {
	products := make([]domain.PurchasedProduct, 0)
	for _, v := range req.PurchasedProducts {
		products = append(products, domain.PurchasedProduct{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		})
	}

	return domain.Order{
		UserID:            req.UserID,
		ShippingAddressID: req.AddressID,
		PaymentMethodID:   req.PaymentMethodID,
		Notes:             req.Notes,
		PurchasedProducts: products,
	}
}

package productpurchaser

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromRequestToOrder(req CreateOrderRequest) domain.Order {
	products := make([]domain.ProductInOrder, 0)
	for _, v := range req.PurchasedProducts {
		products = append(products, domain.ProductInOrder{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		})
	}

	return domain.Order{
		UserID:            req.UserID,
		ShippingAddressID: req.AddressID,
		PaymentMethodID:   req.PaymentMethodID,
		Notes:             req.Notes,
		Products:          products,
	}
}

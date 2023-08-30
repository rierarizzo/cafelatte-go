package productpurchaser

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromRequestToOrder(req OrderCreate) domain.Order {
	products := make([]domain.ProductInOrder, 0)
	for _, v := range req.PurchasedProducts {
		products = append(products, domain.ProductInOrder{
			ProductId: v.ProductId,
			Quantity:  v.Quantity,
		})
	}

	return domain.Order{
		UserId:            req.UserId,
		ShippingAddressId: req.AddressId,
		PaymentMethodId:   req.PaymentMethodId,
		Notes:             req.Notes,
		Products:          products,
	}
}

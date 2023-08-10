package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
)

func requestToPurchasedProduct(req dto.PurchasedProduct) entities.PurchasedProduct {
	return entities.PurchasedProduct{
		ID:        req.ProductID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
}

func RequestToOrder(req dto.PurchaseOrderRequest) entities.PurchaseOrder {
	products := make([]entities.PurchasedProduct, 0)
	for _, v := range req.PurchasedProducts {
		products = append(products, requestToPurchasedProduct(v))
	}

	return entities.PurchaseOrder{
		UserID:            req.UserID,
		ShippingAddressID: req.AddressID,
		PaymentMethodID:   req.PaymentMethodID,
		Notes:             req.Notes,
		PurchasedProducts: products,
	}
}

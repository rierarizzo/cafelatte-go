package productpurchaser

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type DefaultPurchaser struct {
	orderRepository OrderRepository
}

func (purchaser *DefaultPurchaser) Purchase(order domain.Order) (
	int,
	*domain.AppError,
) {
	orderId, appErr := purchaser.orderRepository.InsertPurchaseOrder(order)
	if appErr != nil {
		return 0, appErr
	}

	return orderId, nil
}

func New(orderRepository OrderRepository) *DefaultPurchaser {
	return &DefaultPurchaser{orderRepository}
}

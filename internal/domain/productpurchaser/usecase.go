package productpurchaser

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type DefaultPurchaser struct {
	orderRepository OrderRepository
}

func (p *DefaultPurchaser) Purchase(order domain.Order) (int, *domain.AppError) {
	orderID, appErr := p.orderRepository.InsertPurchaseOrder(order)
	if appErr != nil {
		return 0, appErr
	}

	return orderID, nil
}

func NewDefaultPurchaser(orderRepository OrderRepository) *DefaultPurchaser {
	return &DefaultPurchaser{orderRepository}
}

package order

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
)

type Service struct {
	orderRepo IOrderRepository
}

func (p *Service) CreatePurchaseOrder(order Order) (int, *domain.AppError) {
	orderID, appErr := p.orderRepo.InsertPurchaseOrder(order)
	if appErr != nil {
		return 0, appErr
	}

	return orderID, nil
}

func NewOrderService(orderRepo IOrderRepository) *Service {
	return &Service{orderRepo}
}

package services

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
)

type OrderService struct {
	orderRepo ports.IOrderRepository
}

func (p *OrderService) CreatePurchaseOrder(order entities.PurchaseOrder) (int, *domain.AppError) {
	orderID, appErr := p.orderRepo.InsertPurchaseOrder(order)
	if appErr != nil {
		return 0, appErr
	}

	return orderID, nil
}

func NewOrderService(orderRepo ports.IOrderRepository) *OrderService {
	return &OrderService{orderRepo}
}

package usecases

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
)

type PurchaseUsecase struct {
	orderService ports.IOrderService
}

func (p *PurchaseUsecase) PurchaseOrder(order entities.PurchaseOrder) (int, *domain.AppError) {
	orderID, appErr := p.orderService.CreatePurchaseOrder(order)
	if appErr != nil {
		return 0, appErr
	}

	return orderID, nil
}

func NewPurchaseUsecase(orderService ports.IOrderService) *PurchaseUsecase {
	return &PurchaseUsecase{orderService}
}

package purchase

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/order"
)

type Usecase struct {
	orderService order.IOrderService
}

func (p *Usecase) PurchaseOrder(order order.Order) (int, *domain.AppError) {
	orderID, appErr := p.orderService.CreatePurchaseOrder(order)
	if appErr != nil {
		return 0, appErr
	}

	return orderID, nil
}

func NewPurchaseUsecase(orderService order.IOrderService) *Usecase {
	return &Usecase{orderService}
}

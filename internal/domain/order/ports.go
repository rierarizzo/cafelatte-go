package order

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
)

type IOrderService interface {
	CreatePurchaseOrder(order Order) (int, *domain.AppError)
}

type IOrderRepository interface {
	InsertPurchaseOrder(order Order) (int, *domain.AppError)
}

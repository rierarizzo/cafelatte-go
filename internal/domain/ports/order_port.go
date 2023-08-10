package ports

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
)

type IOrderService interface {
	CreatePurchaseOrder(order entities.PurchaseOrder) (int, *domain.AppError)
}

type IOrderRepository interface {
	InsertPurchaseOrder(order entities.PurchaseOrder) (int, *domain.AppError)
}

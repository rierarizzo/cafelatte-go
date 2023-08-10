package ports

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
)

type IPurchaseUsecase interface {
	PurchaseOrder(order entities.PurchaseOrder) (int, *domain.AppError)
}

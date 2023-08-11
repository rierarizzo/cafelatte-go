package purchase

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/order"
)

type IPurchaseUsecase interface {
	PurchaseOrder(order order.Order) (int, *domain.AppError)
}

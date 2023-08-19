package productpurchaser

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Purchaser interface {
	Purchase(order domain.Order) (int, *domain.AppError)
}

type OrderRepository interface {
	InsertPurchaseOrder(order domain.Order) (int, *domain.AppError)
}

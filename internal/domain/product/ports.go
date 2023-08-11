package product

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/order"
)

type IProductService interface {
	GetProducts() ([]Product, *domain.AppError)
	GetProductsByCategory(categoryCode string) ([]Product, *domain.AppError)
	GetProductCategories() ([]order.ProductCategory, *domain.AppError)
}

type IProductRepository interface {
	SelectProducts() ([]Product, *domain.AppError)
	SelectProductsByCategory(categoryCode string) ([]Product, *domain.AppError)
	SelectProductCategories() ([]order.ProductCategory, *domain.AppError)
}

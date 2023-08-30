package productmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Manager interface {
	GetProducts() ([]domain.Product, *domain.AppError)
	GetProductsByCategory(categoryCode string) ([]domain.Product,
		*domain.AppError)
	GetProductCategories() ([]domain.ProductCategory, *domain.AppError)
}

type ProductRepository interface {
	SelectProducts() ([]domain.Product, *domain.AppError)
	SelectProductsByCategory(categoryCode string) ([]domain.Product,
		*domain.AppError)
	SelectProductCategories() ([]domain.ProductCategory, *domain.AppError)
}

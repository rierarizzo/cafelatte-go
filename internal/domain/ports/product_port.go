package ports

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
)

type IProductService interface {
	GetProducts() ([]entities.Product, *domain.AppError)
	GetProductsByCategory(categoryCode string) ([]entities.Product, *domain.AppError)
	GetProductCategories() ([]entities.ProductCategory, *domain.AppError)
}

type IProductRepository interface {
	SelectProducts() ([]entities.Product, *domain.AppError)
	SelectProductsByCategory(categoryCode string) ([]entities.Product, *domain.AppError)
	SelectProductCategories() ([]entities.ProductCategory, *domain.AppError)
}

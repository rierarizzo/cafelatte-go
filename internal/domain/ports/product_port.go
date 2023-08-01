package ports

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
)

type IProductService interface {
	GetProducts() ([]entities.Product, error)
	GetProductsByCategory(categoryCode string) ([]entities.Product, error)
	GetProductCategories() ([]entities.ProductCategory, error)
}

type IProductRepo interface {
	SelectProducts() ([]entities.Product, error)
	SelectProductsByCategory(categoryCode string) ([]entities.Product, error)
	SelectProductCategories() ([]entities.ProductCategory, error)
}

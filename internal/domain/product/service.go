package product

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/order"
)

type Service struct {
	productRepo IProductRepository
}

func (p Service) GetProducts() ([]Product, *domain.AppError) {
	products, appErr := p.productRepo.SelectProducts()
	if appErr != nil {
		return nil, appErr
	}

	return products, nil
}

func (p Service) GetProductsByCategory(categoryCode string) ([]Product, *domain.AppError) {
	products, appErr := p.productRepo.SelectProductsByCategory(categoryCode)
	if appErr != nil {
		return nil, appErr
	}

	return products, nil
}

func (p Service) GetProductCategories() ([]order.ProductCategory, *domain.AppError) {
	categories, appErr := p.productRepo.SelectProductCategories()
	if appErr != nil {
		return nil, appErr
	}

	return categories, nil
}

func NewProductService(productRepo IProductRepository) *Service {
	return &Service{productRepo}
}

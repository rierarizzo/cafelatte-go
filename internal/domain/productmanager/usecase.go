package productmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type DefaultManager struct {
	productRepository ProductRepository
}

func (m DefaultManager) GetProducts() ([]domain.Product, *domain.AppError) {
	products, appErr := m.productRepository.SelectProducts()
	if appErr != nil {
		return nil, appErr
	}

	return products, nil
}

func (m DefaultManager) GetProductsByCategory(categoryCode string) ([]domain.Product, *domain.AppError) {
	products, appErr := m.productRepository.SelectProductsByCategory(categoryCode)
	if appErr != nil {
		return nil, appErr
	}

	return products, nil
}

func (m DefaultManager) GetProductCategories() ([]domain.ProductCategory, *domain.AppError) {
	categories, appErr := m.productRepository.SelectProductCategories()
	if appErr != nil {
		return nil, appErr
	}

	return categories, nil
}

func New(productRepository ProductRepository) *DefaultManager {
	return &DefaultManager{productRepository}
}

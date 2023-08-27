package productmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type DefaultManager struct {
	productRepository ProductRepository
}

func (manager DefaultManager) GetProducts() ([]domain.Product, *domain.AppError) {
	products, appErr := manager.productRepository.SelectProducts()
	if appErr != nil {
		return nil, appErr
	}

	return products, nil
}

func (manager DefaultManager) GetProductsByCategory(categoryCode string) ([]domain.Product, *domain.AppError) {
	products, appErr := manager.productRepository.SelectProductsByCategory(categoryCode)
	if appErr != nil {
		return nil, appErr
	}

	return products, nil
}

func (manager DefaultManager) GetProductCategories() ([]domain.ProductCategory, *domain.AppError) {
	categories, appErr := manager.productRepository.SelectProductCategories()
	if appErr != nil {
		return nil, appErr
	}

	return categories, nil
}

func New(productRepository ProductRepository) *DefaultManager {
	return &DefaultManager{productRepository}
}

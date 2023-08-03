package services

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
)

type ProductService struct {
	productRepo ports.IProductRepo
}

func (p ProductService) GetProducts() ([]entities.Product, *domain.AppError) {
	products, appErr := p.productRepo.SelectProducts()
	if appErr != nil {
		return nil, appErr
	}

	return products, nil
}

func (p ProductService) GetProductsByCategory(categoryCode string) ([]entities.Product, *domain.AppError) {
	products, appErr := p.productRepo.SelectProductsByCategory(categoryCode)
	if appErr != nil {
		return nil, appErr
	}

	return products, nil
}

func (p ProductService) GetProductCategories() ([]entities.ProductCategory, *domain.AppError) {
	categories, appErr := p.productRepo.SelectProductCategories()
	if appErr != nil {
		return nil, appErr
	}

	return categories, nil
}

func NewProductService(productRepo ports.IProductRepo) *ProductService {
	return &ProductService{productRepo}
}

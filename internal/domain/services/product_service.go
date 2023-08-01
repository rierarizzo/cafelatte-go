package services

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
)

type ProductService struct {
	productRepo ports.IProductRepo
}

func (p ProductService) GetProducts() ([]entities.Product, error) {
	products, err := p.productRepo.SelectProducts()
	if err != nil {
		var appErr *domain.AppError
		converted := errors.As(err, &appErr)
		if !converted {
			return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
		}

		return nil, appErr
	}

	return products, nil
}

func (p ProductService) GetProductsByCategory(categoryCode string) ([]entities.Product, error) {
	products, err := p.productRepo.SelectProductsByCategory(categoryCode)
	if err != nil {
		var appErr *domain.AppError
		converted := errors.As(err, &appErr)
		if !converted {
			return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
		}

		return nil, appErr
	}

	return products, nil
}

func (p ProductService) GetProductCategories() ([]entities.ProductCategory, error) {
	categories, err := p.productRepo.SelectProductCategories()
	if err != nil {
		var appErr *domain.AppError
		converted := errors.As(err, &appErr)
		if !converted {
			return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
		}

		return nil, appErr
	}

	return categories, nil
}

func NewProductService(productRepo ports.IProductRepo) *ProductService {
	return &ProductService{productRepo}
}

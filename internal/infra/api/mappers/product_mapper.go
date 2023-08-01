package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/infra/api/dto"
)

func FromProductToProductResponse(product entities.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		ImageURL:     product.ImageURL,
		Price:        product.Price,
		CategoryCode: product.CategoryCode,
		Stock:        product.Stock,
	}
}

func FromProductCategoryToProductCategoryResponse(category entities.ProductCategory) dto.ProductCategoryResponse {
	return dto.ProductCategoryResponse{
		Code:        category.Code,
		Description: category.Description,
	}
}

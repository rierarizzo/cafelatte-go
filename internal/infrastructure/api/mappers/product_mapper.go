package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
)

func FromProductToProductRes(product entities.Product) dto.ProductResponse {
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

func FromProductSliceToProductResSlice(products []entities.Product) []dto.ProductResponse {
	var res = make([]dto.ProductResponse, 0)
	for _, v := range products {
		res = append(res, FromProductToProductRes(v))
	}

	return res
}

func FromProductCategoryToProductCategoryRes(category entities.ProductCategory) dto.ProductCategoryResponse {
	return dto.ProductCategoryResponse{
		Code:        category.Code,
		Description: category.Description,
	}
}

func FromProductCategorySliceToProductCategoryResSlice(categories []entities.ProductCategory) []dto.ProductCategoryResponse {
	var res = make([]dto.ProductCategoryResponse, 0)
	for _, v := range categories {
		res = append(res, FromProductCategoryToProductCategoryRes(v))
	}

	return res
}

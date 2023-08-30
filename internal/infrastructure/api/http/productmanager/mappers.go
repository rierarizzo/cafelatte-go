package productmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromProductToResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		Id:           product.Id,
		Name:         product.Name,
		Description:  product.Description,
		ImageUrl:     product.ImageUrl,
		Price:        product.Price,
		CategoryCode: product.CategoryCode,
		Stock:        product.Stock,
	}
}

func fromProductsToResponse(products []domain.Product) []ProductResponse {
	var res = make([]ProductResponse, 0)
	for _, v := range products {
		res = append(res, fromProductToResponse(v))
	}

	return res
}

func fromProductCategoryToResponse(category domain.ProductCategory) ProductCategoryResponse {
	return ProductCategoryResponse{
		Code:        category.Code,
		Description: category.Description,
	}
}

func fromProductCategoriesToResponse(categories []domain.ProductCategory) []ProductCategoryResponse {
	var res = make([]ProductCategoryResponse, 0)
	for _, v := range categories {
		res = append(res, fromProductCategoryToResponse(v))
	}

	return res
}

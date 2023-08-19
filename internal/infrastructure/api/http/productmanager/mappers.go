package productmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromProductToResponse(product domain.Product) Response {
	return Response{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		ImageURL:     product.ImageURL,
		Price:        product.Price,
		CategoryCode: product.CategoryCode,
		Stock:        product.Stock,
	}
}

func fromProductsToResponse(products []domain.Product) []Response {
	var res = make([]Response, 0)
	for _, v := range products {
		res = append(res, fromProductToResponse(v))
	}

	return res
}

func fromProductCategoryToResponse(category domain.ProductCategory) CategoryResponse {
	return CategoryResponse{
		Code:        category.Code,
		Description: category.Description,
	}
}

func fromProductCategoriesToResponse(categories []domain.ProductCategory) []CategoryResponse {
	var res = make([]CategoryResponse, 0)
	for _, v := range categories {
		res = append(res, fromProductCategoryToResponse(v))
	}

	return res
}

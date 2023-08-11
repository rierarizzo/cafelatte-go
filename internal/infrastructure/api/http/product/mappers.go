package product

import (
	"github.com/rierarizzo/cafelatte/internal/domain/order"
	"github.com/rierarizzo/cafelatte/internal/domain/product"
)

func FromProductToProductRes(product product.Product) Response {
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

func FromProductSliceToProductResSlice(products []product.Product) []Response {
	var res = make([]Response, 0)
	for _, v := range products {
		res = append(res, FromProductToProductRes(v))
	}

	return res
}

func FromProductCategoryToProductCategoryRes(category order.ProductCategory) CategoryResponse {
	return CategoryResponse{
		Code:        category.Code,
		Description: category.Description,
	}
}

func FromProductCategorySliceToProductCategoryResSlice(categories []order.ProductCategory) []CategoryResponse {
	var res = make([]CategoryResponse, 0)
	for _, v := range categories {
		res = append(res, FromProductCategoryToProductCategoryRes(v))
	}

	return res
}

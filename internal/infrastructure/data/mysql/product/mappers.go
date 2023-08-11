package product

import (
	"github.com/rierarizzo/cafelatte/internal/domain/order"
	"github.com/rierarizzo/cafelatte/internal/domain/product"
	order2 "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/order"
)

func FromProductModelToProduct(model Model) product.Product {
	return product.Product{
		ID:           model.ID,
		Name:         model.Name,
		Description:  model.Description,
		ImageURL:     model.ImageURL,
		Price:        model.Price,
		CategoryCode: model.CategoryCode,
		Stock:        model.Stock,
	}
}

func FromProductModelSliceToProductSlice(productModels []Model) []product.Product {
	var products = make([]product.Product, 0)
	for _, v := range productModels {
		products = append(products, FromProductModelToProduct(v))
	}

	return products
}

func FromProductCategoryModelToProductCategory(model order2.ProductCategoryModel) order.ProductCategory {
	return order.ProductCategory{
		Code:        model.Code,
		Description: model.Description,
	}
}

func FromProductCategoryModelSliceToProductCategorySlice(catModels []order2.ProductCategoryModel) []order.ProductCategory {
	var categories = make([]order.ProductCategory, 0)
	for _, v := range catModels {
		categories = append(categories,
			FromProductCategoryModelToProductCategory(v))
	}

	return categories
}

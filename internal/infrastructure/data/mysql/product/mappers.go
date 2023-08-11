package product

import (
	"github.com/rierarizzo/cafelatte/internal/domain/order"
	"github.com/rierarizzo/cafelatte/internal/domain/product"
)

func fromModelToProduct(model Model) product.Product {
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

func fromModelsToProducts(productModels []Model) []product.Product {
	var products = make([]product.Product, 0)
	for _, v := range productModels {
		products = append(products, fromModelToProduct(v))
	}

	return products
}

func fromCategoryModelToCategory(model CategoryModel) order.ProductCategory {
	return order.ProductCategory{
		Code:        model.Code,
		Description: model.Description,
	}
}

func fromCategoryModelsToCategories(catModels []CategoryModel) []order.ProductCategory {
	var categories = make([]order.ProductCategory, 0)
	for _, v := range catModels {
		categories = append(categories,
			fromCategoryModelToCategory(v))
	}

	return categories
}

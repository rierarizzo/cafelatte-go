package product

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromModelToProduct(model Model) domain.Product {
	return domain.Product{
		Id:           model.Id,
		Name:         model.Name,
		Description:  model.Description,
		ImageUrl:     model.ImageUrl,
		Price:        model.Price,
		CategoryCode: model.CategoryCode,
		Stock:        model.Stock,
	}
}

func fromModelsToProducts(productModels []Model) []domain.Product {
	var products = make([]domain.Product, 0)
	for _, v := range productModels {
		products = append(products, fromModelToProduct(v))
	}

	return products
}

func fromCategoryModelToCategory(model CategoryModel) domain.ProductCategory {
	return domain.ProductCategory{
		Code:        model.Code,
		Description: model.Description,
	}
}

func fromCategoryModelsToCategories(catModels []CategoryModel) []domain.ProductCategory {
	var categories = make([]domain.ProductCategory, 0)
	for _, v := range catModels {
		categories = append(categories,
			fromCategoryModelToCategory(v))
	}

	return categories
}

package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

func FromProductModelToProduct(model models.ProductModel) entities.Product {
	return entities.Product{
		ID:           model.ID,
		Name:         model.Name,
		Description:  model.Description,
		ImageURL:     model.ImageURL,
		Price:        model.Price,
		CategoryCode: model.CategoryCode,
		Stock:        model.Stock,
	}
}

func FromProductModelSliceToProductSlice(productModels []models.ProductModel) []entities.Product {
	var products = make([]entities.Product, 0)
	for _, v := range productModels {
		products = append(products, FromProductModelToProduct(v))
	}

	return products
}

func FromProductCategoryModelToProductCategory(model models.ProductCategoryModel) entities.ProductCategory {
	return entities.ProductCategory{
		Code:        model.Code,
		Description: model.Description,
	}
}

func FromProductCategoryModelSliceToProductCategorySlice(catModels []models.ProductCategoryModel) []entities.ProductCategory {
	var categories = make([]entities.ProductCategory, 0)
	for _, v := range catModels {
		categories = append(categories,
			FromProductCategoryModelToProductCategory(v))
	}

	return categories
}

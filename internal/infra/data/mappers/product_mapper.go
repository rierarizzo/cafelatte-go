package mappers

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/infra/data/models"
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

func FromProductCategoryModelToProductCategory(model models.ProductCategoryModel) entities.ProductCategory {
	return entities.ProductCategory{
		Code:        model.Code,
		Description: model.Description,
	}
}

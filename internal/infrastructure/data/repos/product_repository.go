package repos

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/constants"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
)

type ProductRepository struct {
	db *sqlx.DB
}

var (
	selectProductError         = errors.New("error in selecting product")
	selectProductCategoryError = errors.New("error in selecting product category")
)

func (p ProductRepository) SelectProducts() ([]entities.Product, *domain.AppError) {
	return selectProducts(p.db, "select * from product where Status=true")
}

func (p ProductRepository) SelectProductsByCategory(categoryCode string) ([]entities.Product, *domain.AppError) {
	return selectProducts(p.db,
		"select * from product where CategoryCode=? and Status=true",
		categoryCode)
}

func selectProducts(db *sqlx.DB, query string,
	args ...interface{}) ([]entities.Product, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	var products []entities.Product

	var productsModel []models.ProductModel

	var err error
	if len(args) > 0 {
		err = db.Select(&productsModel, query, args[0])
	} else {
		err = db.Select(&productsModel, query)
	}

	if err != nil {
		log.Errorf("Error in selectProducts: %v", err)
		return nil, domain.NewAppError(selectProductError,
			domain.RepositoryError)
	}

	if productsModel == nil {
		log.Debug("productsModel is empty")
		return products, nil
	}

	for _, v := range productsModel {
		products = append(products, mappers.FromProductModelToProduct(v))
	}

	return products, nil
}

func (p ProductRepository) SelectProductCategories() ([]entities.ProductCategory, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	var productCategories []entities.ProductCategory

	var productCategoriesModel []models.ProductCategoryModel

	err := p.db.Select(&productCategoriesModel, "select * from productcategory")
	if err != nil {
		log.Errorf("Error in SelectProductCategories: %v", err)
		return nil, domain.NewAppError(selectProductCategoryError,
			domain.RepositoryError)
	}

	if productCategoriesModel == nil {
		log.Debug("productCategoriesModel is empty")
		return productCategories, nil
	}

	for _, v := range productCategoriesModel {
		productCategories = append(productCategories,
			mappers.FromProductCategoryModelToProductCategory(v))
	}

	return productCategories, nil
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db}
}

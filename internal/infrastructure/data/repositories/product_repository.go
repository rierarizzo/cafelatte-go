package repositories

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

	var model []models.ProductModel

	var err error
	if len(args) > 0 {
		err = db.Select(&model, query, args[0])
	} else {
		err = db.Select(&model, query)
	}

	if err != nil {
		log.Errorf("Error in selectProducts: %v", err)
		appErr := domain.NewAppError(selectProductError, domain.RepositoryError)
		return nil, appErr
	}

	if model == nil {
		log.Debug("productsModel is empty")
		return []entities.Product{}, nil
	}

	return mappers.FromProductModelSliceToProductSlice(model), nil
}

func (p ProductRepository) SelectProductCategories() ([]entities.ProductCategory, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	var model []models.ProductCategoryModel

	err := p.db.Select(&model, "select * from productcategory")
	if err != nil {
		log.Errorf("Error in SelectProductCategories: %v", err)
		appErr := domain.NewAppError(selectProductCategoryError,
			domain.RepositoryError)
		return nil, appErr
	}

	if model == nil {
		log.Debug("productCategoriesModel is empty")
		return []entities.ProductCategory{}, nil
	}

	return mappers.FromProductCategoryModelSliceToProductCategorySlice(model), nil
}

var (
	selectProductError         = errors.New("select product error")
	selectProductCategoryError = errors.New("select product category error")
)

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db}
}

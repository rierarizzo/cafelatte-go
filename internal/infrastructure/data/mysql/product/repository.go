package product

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	db *sqlx.DB
}

func (p Repository) SelectProducts() ([]domain.Product, *domain.AppError) {
	return selectProducts(p.db, "select * from Product where Status=true")
}

func (p Repository) SelectProductsByCategory(categoryCode string) ([]domain.Product, *domain.AppError) {
	return selectProducts(p.db,
		"select * from Product where CategoryCode=? and Status=true",
		categoryCode)
}

func selectProducts(db *sqlx.DB, query string,
	args ...interface{}) ([]domain.Product, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var model []Model

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
		return []domain.Product{}, nil
	}

	return fromModelsToProducts(model), nil
}

func (p Repository) SelectProductCategories() ([]domain.ProductCategory, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var model []CategoryModel

	err := p.db.Select(&model, "select * from ProductCategory")
	if err != nil {
		log.Errorf("Error in SelectProductCategories: %v", err)
		appErr := domain.NewAppError(selectProductCategoryError,
			domain.RepositoryError)
		return nil, appErr
	}

	if model == nil {
		log.Debug("productCategoriesModel is empty")
		return []domain.ProductCategory{}, nil
	}

	return fromCategoryModelsToCategories(model), nil
}

var (
	selectProductError         = errors.New("select productmanager error")
	selectProductCategoryError = errors.New("select productmanager category error")
)

func NewProductRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}

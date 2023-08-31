package product

import (
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (r Repository) SelectProducts() ([]domain.Product, *domain.AppError) {
	query := `
		SELECT * FROM Product WHERE Status=TRUE
	`
	return selectProducts(r.db, query)
}

func (r Repository) SelectProductsByCategory(categoryCode string) ([]domain.Product, *domain.AppError) {
	query := `
		SELECT * FROM Product WHERE CategoryCode=? AND Status=TRUE
	`
	return selectProducts(r.db, query, categoryCode)
}

func selectProducts(db *sqlx.DB, query string,
	args ...interface{}) ([]domain.Product, *domain.AppError) {
	var model []Model

	var err error
	if len(args) > 0 {
		err = db.Select(&model, query, args[0])
	} else {
		err = db.Select(&model, query)
	}

	if err != nil {
		appErr := domain.NewAppError(err, domain.RepositoryError)
		return nil, appErr
	}

	if model == nil {
		return []domain.Product{}, nil
	}

	products := fromModelsToProducts(model)
	return products, nil
}

func (r Repository) SelectProductCategories() ([]domain.ProductCategory, *domain.AppError) {
	var model []CategoryModel

	query := `
		SELECT * FROM ProductCategory
	`
	err := r.db.Select(&model, query)
	if err != nil {
		appErr := domain.NewAppError(err, domain.RepositoryError)
		return nil, appErr
	}

	if model == nil {
		return []domain.ProductCategory{}, nil
	}

	categories := fromCategoryModelsToCategories(model)
	return categories, nil
}

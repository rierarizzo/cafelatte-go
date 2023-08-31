package address

import (
	"database/sql"
	"errors"
	sqlUtil "github.com/rierarizzo/cafelatte/pkg/utils/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (r Repository) SelectAddressesByUserId(userId int) ([]domain.Address, *domain.AppError) {
	var addressesModel []Model

	query := `
		SELECT * FROM UserAddress WHERE UserId=? AND Status=TRUE
	`
	err := r.db.Select(&addressesModel, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Address{}, nil
		}

		appErr := domain.NewAppError(err, domain.RepositoryError)
		return nil, appErr
	}

	addresses := fromModelsToAddresses(addressesModel)
	return addresses, nil
}

func (r Repository) InsertUserAddress(userId int,
	address domain.Address) (*domain.Address, *domain.AppError) {
	tx, appErr := sqlUtil.StartTransaction(r.db)
	if appErr != nil {
		return nil, appErr
	}

	defer sqlUtil.RollbackIfPanic(tx)

	model := fromAddressToModel(address)

	query := `
		INSERT INTO UserAddress (Type, UserId, ProvinceId, CityId, PostalCode, Detail) 
		VALUES (?,?,?,?,?,?)
	`
	result, appErr := sqlUtil.ExecWithTransaction(tx,
		query,
		model.Type,
		userId,
		model.ProvinceId,
		model.CityId,
		model.PostalCode,
		model.Detail)
	if appErr != nil {
		return nil, appErr
	}

	addressId, appErr := sqlUtil.GetLastInsertedId(result)
	if appErr != nil {
		return nil, appErr
	}
	address.Id = addressId

	if appErr = sqlUtil.CommitTransaction(tx); appErr != nil {
		return nil, appErr
	}

	return &address, nil
}

func (r Repository) SelectCityNameById(id int) (string, *domain.AppError) {
	var cityName string

	query := `
		SELECT Name FROM City WHERE Id=?
	`
	err := r.db.Get(&cityName, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			appErr := domain.NewAppError("city not found", domain.NotFoundError)
			return "", appErr
		}

		appErr := domain.NewAppError(err, domain.RepositoryError)
		return "", appErr
	}

	return cityName, nil
}

func (r Repository) SelectProvinceNameById(id int) (string, *domain.AppError) {
	var provinceName string

	query := `
		SELECT Name FROM Province WHERE Id=?
	`
	err := r.db.Get(&provinceName, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			appErr := domain.NewAppError("province not found",
				domain.NotFoundError)
			return "", appErr
		}

		appErr := domain.NewAppError(err, domain.RepositoryError)
		return "", appErr
	}

	return provinceName, nil
}

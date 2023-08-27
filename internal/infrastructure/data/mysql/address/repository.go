package address

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
)

var (
	selectAddressError = errors.New("select addressmanager error")
	insertAddressError = errors.New("insert addressmanager error")
)

type Repository struct {
	db *sqlx.DB
}

func (repository Repository) SelectAddressesByUserId(userId int) ([]domain.Address, *domain.AppError) {
	var addressesModel []Model

	var query = "select * from UserAddress where UserID=? and Status=true"

	err := repository.db.Select(&addressesModel, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			appErr := domain.NewAppErrorWithType(domain.NotFoundError)
			return nil, appErr
		}

		appErr := domain.NewAppError(selectAddressError, domain.RepositoryError)
		return nil, appErr
	}

	return fromModelsToAddresses(addressesModel), nil
}

func (repository Repository) InsertUserAddress(userId int,
	address domain.Address) (*domain.Address, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	rollbackAndError := func(tx *sqlx.Tx, err error) *domain.AppError {
		_ = tx.Rollback()

		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return domain.NewAppError(insertAddressError, domain.RepositoryError)
	}

	tx, err := repository.db.Beginx()
	if err != nil {
		return nil, rollbackAndError(tx, err)
	}

	addressModel := fromAddressToModel(address)
	result, err := tx.Exec(`insert into UserAddress (
                         Type, 
                         UserID, 
                         ProvinceID, 
                         CityID, 
                         PostalCode, 
                         Detail
                ) values (?,?,?,?,?,?)`, addressModel.Type, userId,
		addressModel.ProvinceID, addressModel.CityID, addressModel.PostalCode,
		addressModel.Detail)
	if err != nil {
		return nil, rollbackAndError(tx, err)
	}

	addressID, _ := result.LastInsertId()
	address.ID = int(addressID)

	err = tx.Commit()
	if err != nil {
		return nil, rollbackAndError(tx, err)
	}

	return &address, nil
}

func (repository Repository) SelectCityNameById(id int) (string, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var cityName string
	var query = "select Name from City where ID=?"

	err := repository.db.Get(&cityName, query, id)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return "", domain.NewAppErrorWithType(domain.NotFoundError)
		}

		appErr := domain.NewAppError(selectAddressError, domain.RepositoryError)
		return "", appErr
	}

	return cityName, nil
}

func (repository Repository) SelectProvinceNameById(id int) (string, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var provinceName string
	var query = "select Name from Province where ID=?"

	err := repository.db.Get(&provinceName, query, id)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return "", domain.NewAppErrorWithType(domain.NotFoundError)
		}

		appErr := domain.NewAppError(selectAddressError, domain.RepositoryError)
		return "", appErr
	}

	return provinceName, nil
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}

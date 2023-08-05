package repos

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/constants"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
	"sync"
)

type AddressRepository struct {
	db *sqlx.DB
}

func (r AddressRepository) SelectAddressesByUserID(userID int) ([]entities.Address, *domain.AppError) {
	var addressesModel []models.AddressModel

	var query = "select * from useraddress where UserID=? and Status=true"

	err := r.db.Select(&addressesModel, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			appErr := domain.NewAppErrorWithType(domain.NotFoundError)
			return nil, appErr
		}

		appErr := domain.NewAppError(selectAddressError, domain.RepositoryError)
		return nil, appErr
	}

	return mappers.ModelSliceToAddresses(addressesModel), nil
}

func (r AddressRepository) InsertUserAddresses(userID int,
	addresses []entities.Address) ([]entities.Address, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	returnError := func(err error) *domain.AppError {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return domain.NewAppError(insertAddressError, domain.RepositoryError)
	}

	tx, err := r.db.Begin()
	if err != nil {
		return nil, returnError(err)
	}

	insertStmnt, err := tx.Prepare(`insert into useraddress (
                         Type, 
                         UserID, 
                         ProvinceID, 
                         CityID, 
                         PostalCode, 
                         Detail
                ) values (?,?,?,?,?,?)`)
	if err != nil {
		return nil, returnError(err)
	}

	sem := make(chan struct{}, 5)

	errCh := make(chan error, len(addresses))
	var wg sync.WaitGroup

	for _, v := range addresses {
		wg.Add(1)
		sem <- struct{}{}

		go func(address entities.Address) {
			defer func() {
				wg.Done()
				<-sem
			}()
			addressModel := mappers.AddressToModel(address)

			result, err := insertStmnt.Exec(addressModel.Type, userID,
				addressModel.ProvinceID, addressModel.CityID,
				addressModel.PostalCode, addressModel.Detail)
			if err != nil {
				errCh <- err
				return
			}

			addressID, _ := result.LastInsertId()
			address.ID = int(addressID)
		}(v)

	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		_ = tx.Rollback()
		return nil, returnError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, returnError(err)
	}

	return addresses, nil
}

func (r AddressRepository) SelectCityNameByCityID(cityID int) (string, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	var cityName string
	var query = "select Name from city where ID=?"

	err := r.db.Get(&cityName, query, cityID)
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

func (r AddressRepository) SelectProvinceNameByProvinceID(cityID int) (string, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	var provinceName string
	var query = "select Name from province where ID=?"

	err := r.db.Get(&provinceName, query, cityID)
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

var (
	selectAddressError = errors.New("select address error")
	insertAddressError = errors.New("insert address error")
)

func NewAddressRepository(db *sqlx.DB) *AddressRepository {
	return &AddressRepository{db}
}

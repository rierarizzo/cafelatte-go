package address

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	selectAddressError = errors.New("select addressmanager error")
	insertAddressError = errors.New("insert addressmanager error")
)

type Repository struct {
	db *sqlx.DB
}

func (r Repository) SelectAddressesByUserID(userID int) ([]domain.Address, *domain.AppError) {
	var addressesModel []Model

	var query = "select * from UserAddress where UserID=? and Status=true"

	err := r.db.Select(&addressesModel, query, userID)
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

func (r Repository) InsertUserAddresses(userID int,
	addresses []domain.Address) ([]domain.Address, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	rollbackAndError := func(tx *sqlx.Tx, err error) *domain.AppError {
		_ = tx.Rollback()

		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return domain.NewAppError(insertAddressError, domain.RepositoryError)
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, rollbackAndError(tx, err)
	}

	insertStmnt, err := tx.Prepare(`insert into UserAddress (
                         Type, 
                         UserID, 
                         ProvinceID, 
                         CityID, 
                         PostalCode, 
                         Detail
                ) values (?,?,?,?,?,?)`)
	if err != nil {
		return nil, rollbackAndError(tx, err)
	}

	sem := make(chan struct{}, 5)

	errCh := make(chan error, len(addresses))
	var wg sync.WaitGroup

	for _, v := range addresses {
		wg.Add(1)
		sem <- struct{}{}

		go func(address domain.Address) {
			defer func() {
				wg.Done()
				<-sem
			}()
			addressModel := fromAddressToModel(address)

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
		return nil, rollbackAndError(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, rollbackAndError(tx, err)
	}

	return addresses, nil
}

func (r Repository) SelectCityNameByCityID(cityID int) (string, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var cityName string
	var query = "select Name from City where ID=?"

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

func (r Repository) SelectProvinceNameByProvinceID(cityID int) (string, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var provinceName string
	var query = "select Name from Province where ID=?"

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

func NewAddressRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}

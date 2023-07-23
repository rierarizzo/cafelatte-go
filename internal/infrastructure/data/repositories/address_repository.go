package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
	"sync"
)

type AddressRepository struct {
	db *sqlx.DB
}

func (a AddressRepository) SelectAddressByID(userID int, addressID int) (*entities.Address, error) {
	var addressModel models.AddressModel

	query := "select * from useraddress where ID=? and UserID=? and Status=true"
	err := a.db.Get(&addressModel, query, userID, addressID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WrapError(errors.ErrRecordNotFound, err.Error())
		}
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	return mappers.FromAddressModelToAddress(addressModel), nil
}

func (a AddressRepository) SelectAddressesByUserID(userID int) ([]entities.Address, error) {
	var addressesModel []models.AddressModel

	query := "select * from useraddress where UserID=? and Status=true"
	err := a.db.Select(&addressesModel, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WrapError(errors.ErrRecordNotFound, err.Error())
		}
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	var addresses []entities.Address
	for _, v := range addressesModel {
		addresses = append(addresses, *mappers.FromAddressModelToAddress(v))
	}

	return addresses, nil
}

func (a AddressRepository) SelectCityNameByCityID(cityID int) (string, error) {
	var cityName string

	query := "select Name from city where ID=?"
	err := a.db.Get(&cityName, query, cityID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.WrapError(errors.ErrRecordNotFound, err.Error())
		}
		return "", errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	return cityName, nil
}

func (a AddressRepository) SelectProvinceNameByProvinceID(cityID int) (string, error) {
	var provinceName string

	query := "select Name from province where ID=?"
	err := a.db.Get(&provinceName, query, cityID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.WrapError(errors.ErrRecordNotFound, err.Error())
		}
		return "", errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	return provinceName, nil
}

func (a AddressRepository) InsertUserAddresses(userID int, addresses []entities.Address) ([]entities.Address, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	insertStmnt, err := tx.Prepare(
		`insert into useraddress (Type, UserID, ProvinceID, CityID, PostalCode, Detail) 
			values (?,?,?,?,?,?)`)
	if err != nil {
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
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
			addressModel := mappers.FromAddressToAddressModel(address)

			result, err := insertStmnt.Exec(addressModel.Type, userID, addressModel.ProvinceID,
				addressModel.CityID, addressModel.PostalCode, addressModel.Detail)
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
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	return addresses, nil
}

func NewAddressRepository(db *sqlx.DB) *AddressRepository {
	return &AddressRepository{db}
}

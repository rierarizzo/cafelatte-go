package repos

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infra/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infra/data/models"
	"sync"
)

type AddressRepo struct {
	db *sqlx.DB
}

var (
	selectAddressError = errors.New("errors in selecting address(es)")
	insertAddressError = errors.New("errors in inserting address")
)

func (r AddressRepo) SelectAddressesByUserID(userID int) (
	[]entities.Address,
	error,
) {
	var addressesModel []models.AddressModel

	query := "select * from useraddress where UserID=? and Status=true"
	err := r.db.Select(&addressesModel, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NewAppErrorWithType(domain.NotFoundError)
		}
		return nil, domain.NewAppError(
			errors.Join(selectAddressError, err),
			domain.RepositoryError,
		)
	}

	var addresses []entities.Address
	for _, v := range addressesModel {
		addresses = append(addresses, *mappers.FromAddressModelToAddress(v))
	}

	return addresses, nil
}

func (r AddressRepo) InsertUserAddresses(
	userID int,
	addresses []entities.Address,
) ([]entities.Address, error) {
	returnRepoError := func(err error) error {
		return domain.NewAppError(
			errors.Join(insertAddressError, err),
			domain.RepositoryError,
		)
	}

	tx, err := r.db.Begin()
	if err != nil {
		return nil, returnRepoError(err)
	}

	insertStmnt, err := tx.Prepare(
		`insert into useraddress (
                         Type, 
                         UserID, 
                         ProvinceID, 
                         CityID, 
                         PostalCode, 
                         Detail
                ) values (?,?,?,?,?,?)`,
	)
	if err != nil {
		return nil, returnRepoError(err)
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

			result, err := insertStmnt.Exec(
				addressModel.Type,
				userID,
				addressModel.ProvinceID,
				addressModel.CityID,
				addressModel.PostalCode,
				addressModel.Detail,
			)
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
		return nil, returnRepoError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, returnRepoError(err)
	}

	return addresses, nil
}

func (r AddressRepo) SelectCityNameByCityID(cityID int) (string, error) {
	var cityName string

	query := "select Name from city where ID=?"
	err := r.db.Get(&cityName, query, cityID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", domain.NewAppErrorWithType(domain.NotFoundError)
		}
		return "", domain.NewAppError(
			errors.Join(selectAddressError, err),
			domain.RepositoryError,
		)
	}

	return cityName, nil
}

func (r AddressRepo) SelectProvinceNameByProvinceID(cityID int) (
	string,
	error,
) {
	var provinceName string

	query := "select Name from province where ID=?"
	err := r.db.Get(&provinceName, query, cityID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", domain.NewAppErrorWithType(domain.NotFoundError)
		}
		return "", domain.NewAppError(
			errors.Join(selectAddressError, err),
			domain.RepositoryError,
		)
	}

	return provinceName, nil
}

func NewAddressRepository(db *sqlx.DB) *AddressRepo {
	return &AddressRepo{db}
}

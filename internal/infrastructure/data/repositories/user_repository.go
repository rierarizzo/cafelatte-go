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

type UserRepository struct {
	db *sqlx.DB
}

const (
	concurrencyLimit             = 5
	selectAddressesByUserIDQuery = "select * from UserAddress ua where ua.UserID=? and u.Status=true"
	selectCardsByUserIDQuery     = "select * from UserPaymentCard upc where upc.UserID=? and u.Status=true"
)

func (ur *UserRepository) SelectAllUsers() ([]entities.User, error) {
	var usersModel []models.UserModel

	query := "select * from User u where u.Status=true"
	err := ur.db.Select(&usersModel, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []entities.User{}, nil
		} else {
			return nil, handleSQLError(err)
		}
	}

	sem := make(chan struct{}, concurrencyLimit)

	errCh := make(chan error, len(usersModel))
	var wg sync.WaitGroup

	for i, v := range usersModel {
		wg.Add(1)
		sem <- struct{}{}

		go func(userIndex int, user models.UserModel) {
			defer func() {
				wg.Done()
				<-sem
			}()

			var addressesModel []models.AddressModel
			var cardsModel []models.PaymentCardModel

			err := ur.db.Select(&addressesModel, selectAddressesByUserIDQuery, user.ID)
			if err != nil && err != sql.ErrNoRows {
				errCh <- err
				return
			}

			err = ur.db.Select(&cardsModel, selectCardsByUserIDQuery, user.ID)
			if err != nil && err != sql.ErrNoRows {
				errCh <- err
				return
			}

			usersModel[userIndex].Addresses = addressesModel
			usersModel[userIndex].PaymentCards = cardsModel
		}(i, v)

	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		return nil, handleSQLError(err)
	}

	var users []entities.User
	for _, k := range usersModel {
		users = append(users, *mappers.FromUserModelToUser(k))
	}

	return users, nil
}

func (ur *UserRepository) SelectUserByID(userID int) (*entities.User, error) {
	var userModel models.UserModel

	err := ur.db.Get(&userModel, "select * from User u where u.ID=? and u.Status=true", userID)
	if err != nil {
		return nil, handleSQLError(err)
	}

	var addressesModel []models.AddressModel
	var cardsModel []models.PaymentCardModel

	err = ur.db.Select(&addressesModel, selectAddressesByUserIDQuery, userModel.ID)
	if err != nil && err != sql.ErrNoRows {
		return nil, handleSQLError(err)
	}

	err = ur.db.Select(&cardsModel, selectCardsByUserIDQuery, userModel.ID)
	if err != nil && err != sql.ErrNoRows {
		return nil, handleSQLError(err)
	}

	userModel.Addresses = addressesModel
	userModel.PaymentCards = cardsModel

	return mappers.FromUserModelToUser(userModel), nil
}

func (ur *UserRepository) SelectUserByEmail(email string) (*entities.User, error) {
	var userModel models.UserModel

	err := ur.db.Get(&userModel, "select * from User u where u.Email=? and u.Status=true", email)
	if err != nil {
		return nil, handleSQLError(err)
	}

	var addressesModel []models.AddressModel
	var cardsModel []models.PaymentCardModel

	err = ur.db.Select(&addressesModel, selectAddressesByUserIDQuery, userModel.ID)
	if err != nil && err != sql.ErrNoRows {
		return nil, handleSQLError(err)
	}

	err = ur.db.Select(&cardsModel, selectCardsByUserIDQuery, userModel.ID)
	if err != nil && err != sql.ErrNoRows {
		return nil, handleSQLError(err)
	}

	userModel.Addresses = addressesModel
	userModel.PaymentCards = cardsModel

	return mappers.FromUserModelToUser(userModel), nil
}

func (ur *UserRepository) InsertUser(user entities.User) (*entities.User, error) {
	userModel := mappers.FromUserToUserModel(user)

	result, err := ur.db.Exec(
		`insert into User (Username, Name, Surname, PhoneNumber, Email, Password, RoleCode) 
			values (?,?,?,?,?,?,?)`,
		userModel.Username, userModel.Name, userModel.Surname, userModel.PhoneNumber,
		userModel.Email, userModel.Password, userModel.RoleCode)
	if err != nil {
		return nil, handleSQLError(err)
	}

	lastUserID, _ := result.LastInsertId()

	userModel.ID = int(lastUserID)
	return mappers.FromUserModelToUser(*userModel), nil
}

func (ur *UserRepository) InsertUserPaymentCards(userID int, cards []entities.PaymentCard) ([]entities.PaymentCard, error) {
	tx, err := ur.db.Begin()
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	insertStmnt, err := tx.Prepare(
		`insert into UserPaymentCard (Type, UserID, Company, HolderName, Number, ExpirationDate, CVV) 
			values (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	sem := make(chan struct{}, concurrencyLimit)

	errCh := make(chan error, len(cards))
	var wg sync.WaitGroup

	for _, v := range cards {
		wg.Add(1)
		sem <- struct{}{}

		go func(card entities.PaymentCard) {
			defer func() {
				wg.Done()
				<-sem
			}()
			cardModel := mappers.FromPaymentCardToPaymentCardModel(card)

			result, err := insertStmnt.Exec(cardModel.Type, userID, cardModel.Company,
				cardModel.HolderName, cardModel.Number, cardModel.ExpirationDate, cardModel.CVV)
			if err != nil {
				errCh <- err
				return
			}

			cardID, _ := result.LastInsertId()
			card.ID = int(cardID)
		}(v)
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		_ = tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	return cards, nil
}

func (ur *UserRepository) InsertUserAddresses(userID int, addresses []entities.Address) ([]entities.Address, error) {
	tx, err := ur.db.Begin()
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	insertStmnt, err := tx.Prepare(
		`insert into UserAddress (Type, UserID, ProvinceID, CityID, PostalCode, Detail) 
			values (?,?,?,?,?,?)`)
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	sem := make(chan struct{}, concurrencyLimit)

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
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	return addresses, nil
}

func (ur *UserRepository) UpdateUser(userID int, user entities.User) error {
	userModel := mappers.FromUserToUserModel(user)

	query := "update User set Username=?, Name=?, Surname=?, PhoneNumber=? where ID=?"

	_, err := ur.db.Exec(query, userModel.Name, userModel.Surname, userModel.PhoneNumber, userID)
	if err != nil {
		return handleSQLError(err)
	}

	return nil
}

func handleSQLError(sqlError error) error {
	switch sqlError {
	case sql.ErrNoRows:
		return errors.ErrRecordNotFound
	default:
		return errors.ErrUnexpected
	}
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

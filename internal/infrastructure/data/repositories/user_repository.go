package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
	"sync"
)

type UserRepository struct {
	db *sqlx.DB
}

func (ur *UserRepository) SelectAllUsers() ([]entities.User, error) {
	var userModel []models.UserModel

	query := "SELECT * FROM user"
	err := ur.db.Select(&userModel, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []entities.User{}, nil
		} else {
			return nil, handleSQLError(err)
		}
	}

	var users []entities.User
	for _, k := range userModel {
		users = append(users, *k.ToUserCore())
	}

	return users, nil
}

func (ur *UserRepository) SelectUserById(userID int) (*entities.User, error) {
	var userModel models.UserModel

	query := "SELECT * FROM user u WHERE u.id=?"
	err := ur.db.Get(&userModel, query, userID)
	if err != nil {
		return nil, handleSQLError(err)
	}

	return userModel.ToUserCore(), nil
}

func (ur *UserRepository) SelectUserByEmail(email string) (*entities.User, error) {
	var userModel models.UserModel

	query := "SELECT * FROM user u WHERE u.email=?"
	err := ur.db.Get(&userModel, query, email)
	if err != nil {
		return nil, handleSQLError(err)
	}

	return userModel.ToUserCore(), nil
}

func (ur *UserRepository) InsertUser(user entities.User) (*entities.User, error) {
	var userModel models.UserModel
	userModel.LoadFromUserCore(user)

	result, err := ur.db.Exec(
		`INSERT INTO User (Username, Name, Surname, PhoneNumber, Email, Password, RoleCode) 
			VALUES (?,?,?,?,?,?,?)`,
		userModel.Username, userModel.Name, userModel.Surname, userModel.PhoneNumber,
		userModel.Email, userModel.Password, userModel.RoleCode)
	if err != nil {
		return nil, handleSQLError(err)
	}

	lastUserID, _ := result.LastInsertId()

	userModel.ID = int(lastUserID)
	return userModel.ToUserCore(), nil
}

func (ur *UserRepository) InsertUserPaymentCards(userID int, cards []entities.PaymentCard) ([]entities.PaymentCard, error) {
	tx, err := ur.db.Begin()
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	insertStmnt, err := tx.Prepare(
		`INSERT INTO UserPaymentCard (Type, UserID, Company, Issuer, HolderName, Number, ExpirationDate, CVV) 
			VALUES (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	errCh := make(chan error, len(cards))
	var wg sync.WaitGroup

	for _, v := range cards {
		wg.Add(1)

		go func(card entities.PaymentCard) {
			defer wg.Done()

			var cardModel models.PaymentCardModel
			cardModel.LoadFromPaymentCardCore(card)
			cardModel.UserID = userID

			result, err := insertStmnt.Exec(cardModel.Type, cardModel.UserID, cardModel.Company, cardModel.Issuer,
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
		`INSERT INTO UserAddress (Type, UserID, ProvinceID, CityID, PostalCode, Detail) 
			VALUES (?,?,?,?,?,?)`)
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	for _, v := range addresses {
		var addressModel models.AddressModel
		addressModel.LoadFromAddressCore(v)
		addressModel.UserID = userID

		result, err := insertStmnt.Exec(addressModel.Type, addressModel.UserID, addressModel.ProvinceID,
			addressModel.CityID, addressModel.PostalCode, addressModel.Detail)
		if err != nil {
			_ = tx.Rollback()
			return nil, errors.ErrUnexpected
		}

		addressID, _ := result.LastInsertId()
		v.ID = int(addressID)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	return addresses, nil
}

func (ur *UserRepository) UpdateUser(userID int, user entities.User) error {
	var userModel models.UserModel
	userModel.LoadFromUserCore(user)

	query := "UPDATE user SET name=?, surname=?, phone_number=?, email=?, password=? WHERE id=?"

	_, err := ur.db.Exec(query, user.Name, user.Surname, user.PhoneNumber, user.Email, user.Password, userID)
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

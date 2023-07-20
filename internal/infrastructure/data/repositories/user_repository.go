package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func (ur *UserRepository) GetAllUsers() ([]entities.User, error) {
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

func (ur *UserRepository) GetUserById(userID int) (*entities.User, error) {
	var userModel models.UserModel

	query := "SELECT * FROM user u WHERE u.id=?"
	err := ur.db.Get(&userModel, query, userID)
	if err != nil {
		return nil, handleSQLError(err)
	}

	return userModel.ToUserCore(), nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var userModel models.UserModel

	query := "SELECT * FROM user u WHERE u.email=?"
	err := ur.db.Get(&userModel, query, email)
	if err != nil {
		return nil, handleSQLError(err)
	}

	return userModel.ToUserCore(), nil
}

func (ur *UserRepository) CreateUser(user entities.User) (*entities.User, error) {
	var userModel models.UserModel
	userModel.LoadFromUserCore(user)

	handleTXError := func(tx *sql.Tx, err error) (*entities.User, error) {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return nil, errors.ErrUnexpected
		}

		return nil, handleSQLError(err)
	}

	tx, err := ur.db.Begin()
	if err != nil {
		return nil, handleSQLError(err)
	}

	r, err := tx.Exec(
		`INSERT INTO User (Username, Name, Surname, PhoneNumber, Email, Password, RoleCode) 
			VALUES (?,?,?,?,?,?,?)`,
		userModel.Username, userModel.Name, userModel.Surname, userModel.PhoneNumber,
		userModel.Email, userModel.Password, userModel.RoleCode)
	if err != nil {
		return handleTXError(tx, err)
	}

	lastUserID, _ := r.LastInsertId()

	addressStmt, err := tx.Prepare(
		`INSERT INTO UserAddress (Type, UserID, ProvinceID, CityID, PostalCode, Detail) 
			VALUES (?,?,?,?,?,?)`)
	if err != nil {
		return handleTXError(tx, err)
	}

	for _, v := range user.Addresses {
		var addressModel models.AddressModel
		addressModel.LoadFromAddressCore(v)
		addressModel.UserID = int(lastUserID)

		_, err = addressStmt.Exec(addressModel.Type, addressModel.UserID, addressModel.ProvinceID,
			addressModel.CityID, addressModel.PostalCode, addressModel.Detail)
		if err != nil {
			return handleTXError(tx, err)
		}
	}
	_ = addressStmt.Close()

	paymentCardStmt, err := tx.Prepare(
		`INSERT INTO UserPaymentCard (Type, UserID, Company, Issuer, HolderName, Number, ExpirationDate, CVV) 
			VALUES (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return handleTXError(tx, err)
	}

	for _, v := range user.PaymentCards {
		var cardModel models.PaymentCardModel
		cardModel.LoadFromPaymentCardCore(v)
		cardModel.UserID = int(lastUserID)
		_, err = paymentCardStmt.Exec(cardModel.Type, cardModel.UserID, cardModel.Company, cardModel.Issuer,
			cardModel.HolderName, cardModel.Number, cardModel.ExpirationDate, cardModel.CVV)
		if err != nil {
			return handleTXError(tx, err)
		}
	}
	_ = paymentCardStmt.Close()

	err = tx.Commit()
	if err != nil {
		return handleTXError(tx, err)
	}

	userModel.ID = int(lastUserID)
	return userModel.ToUserCore(), nil
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

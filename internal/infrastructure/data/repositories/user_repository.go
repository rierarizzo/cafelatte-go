package repositories

import (
	"database/sql"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
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

	tx, err := ur.db.Begin()
	if err != nil {
		return nil, handleSQLError(err)
	}
	defer func(tx *sql.Tx) {
		err = tx.Rollback()
		if err != nil {
			logrus.Error("error while rollback")
		}
	}(tx)

	r, err := tx.Exec(
		`INSERT INTO User (Username, Name, Surname, PhoneNumber, Email, Password, RoleCode) 
			VALUES (?,?,?,?,?,?,?)`,
		userModel.Username, userModel.Name, userModel.Surname, userModel.PhoneNumber,
		userModel.Email, userModel.Password, userModel.RoleCode)
	if err != nil {
		return nil, handleSQLError(err)
	}

	userLastID, _ := r.LastInsertId()

	var addressesModel []models.AddressModel
	for _, v := range user.Addresses {
		var addressModel models.AddressModel
		addressModel.LoadFromAddressCore(v)
		addressModel.UserID = int(userLastID)
		addressesModel = append(addressesModel, addressModel)
	}

	for _, v := range addressesModel {
		_, err = tx.Exec(
			`INSERT INTO UserAddress (Type, UserID, ProvinceID, CityID, PostalCode, Detail) 
				VALUES (?,?,?,?,?,?)`,
			v.Type, v.UserID, v.ProvinceID, v.CityID, v.PostalCode, v.Detail)
		if err != nil {
			return nil, handleSQLError(err)
		}
	}

	var paymentCardsModel []models.PaymentCardModel
	for _, v := range user.PaymentCards {
		var cardModel models.PaymentCardModel
		cardModel.LoadFromPaymentCardCore(v)
		cardModel.UserID = int(userLastID)
		paymentCardsModel = append(paymentCardsModel, cardModel)
	}

	for _, v := range paymentCardsModel {
		_, err = tx.Exec(
			`INSERT INTO UserPaymentCard (Type, UserID, Company, Issuer, HolderName, Number, ExpirationDate, CVV) 
				VALUES (?,?,?,?,?,?,?,?)`,
			v.Type, v.UserID, v.Company, v.Issuer, v.HolderName, v.Number, v.ExpirationDate, v.CVV)
		if err != nil {
			return nil, handleSQLError(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, handleSQLError(err)
	}

	userModel.ID = int(userLastID)
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

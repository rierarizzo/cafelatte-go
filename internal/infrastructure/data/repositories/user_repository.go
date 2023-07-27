package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

type UserRepository struct {
	db *sqlx.DB
}

const selectUserWithAllFieldsQuery = `select 
    			u.ID as 'UserID',
    			u.Username as 'UserUsername',
				u.Name as 'UserName',
				u.Surname as 'UserSurname',
				u.PhoneNumber as 'UserPhoneNumber',
				u.Email as 'UserEmail',
				u.Password as 'UserPassword',
				u.RoleCode as 'UserRoleCode',
				u.Status as 'UserStatus',
				u.CreatedAt as 'UserCreatedAt',
				u.UpdatedAt as 'UserUpdatedAt',
				ua.ID as 'AddressID',
				ua.Type as 'AddressType',
				ua.ProvinceID as 'AddressProvinceID',
				ua.CityID as 'AddressCityID',
				ua.PostalCode as 'AddressPostalCode',
				ua.Detail as 'AddressDetail',
				ua.Status as 'AddressStatus',
				ua.CreatedAt as 'AddressCreatedAt',
				ua.UpdatedAt as 'AddressUpdatedAt',
				up.ID as 'CardID',
				up.Type as 'CardType',
				up.Company as 'CardCompany',
				up.HolderName as 'CardHolderName',
				up.Number as 'CardNumber',
				up.ExpirationYear as 'CardExpirationYear',
				up.ExpirationMonth as 'CardExpirationMonth',
				up.CVV as 'CardCVV',
				up.Status as 'CardStatus',
				up.CreatedAt as 'CardCreatedAt',
				up.UpdatedAt as 'CardUpdatedAt'
			from user u inner join useraddress ua on u.ID = ua.UserID inner join userpaymentcard up
    		on u.ID = up.UserID where u.Status=true and ua.Status=true and up.Status=true`

func (ur *UserRepository) SelectAllUsers() ([]entities.User, error) {
	var temporaryUsers []models.TemporaryUserModel

	err := ur.db.Select(&temporaryUsers, selectUserWithAllFieldsQuery)
	if err != nil {
		if err == sql.ErrNoRows {
			return []entities.User{}, nil
		} else {
			return nil, err
		}
	}

	return mappers.FromTemporaryUsersModelToUserSlice(temporaryUsers), nil
}

func (ur *UserRepository) SelectUserByID(userID int) (*entities.User, error) {
	var temporaryUsers []models.TemporaryUserModel

	err := ur.db.Select(&temporaryUsers, selectUserWithAllFieldsQuery+" and u.ID=?", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WrapError(errors.ErrRecordNotFound, err.Error())
		} else {
			return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
		}
	}

	users := mappers.FromTemporaryUsersModelToUserSlice(temporaryUsers)
	return &users[0], nil
}

func (ur *UserRepository) SelectUserByEmail(email string) (*entities.User, error) {
	var temporaryUsers []models.TemporaryUserModel

	err := ur.db.Select(&temporaryUsers, selectUserWithAllFieldsQuery+" and u.Email=?", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WrapError(errors.ErrRecordNotFound, err.Error())
		} else {
			return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
		}
	}

	users := mappers.FromTemporaryUsersModelToUserSlice(temporaryUsers)
	return &users[0], nil
}

func (ur *UserRepository) InsertUser(user entities.User) (*entities.User, error) {
	userModel := mappers.FromUserToUserModel(user)

	result, err := ur.db.Exec(
		`insert into user (Username, Name, Surname, PhoneNumber, Email, Password, RoleCode) 
			values (?,?,?,?,?,?,?)`,
		userModel.Username, userModel.Name, userModel.Surname, userModel.PhoneNumber,
		userModel.Email, userModel.Password, userModel.RoleCode)
	if err != nil {
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	lastUserID, _ := result.LastInsertId()

	userModel.ID = int(lastUserID)
	return mappers.FromUserModelToUser(*userModel), nil
}

func (ur *UserRepository) UpdateUser(userID int, user entities.User) error {
	userModel := mappers.FromUserToUserModel(user)

	query := "update user set Username=?, Name=?, Surname=?, PhoneNumber=? where ID=?"

	_, err := ur.db.Exec(query, userModel.Name, userModel.Surname, userModel.PhoneNumber, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.WrapError(errors.ErrRecordNotFound, err.Error())
		}
		return errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	return nil
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

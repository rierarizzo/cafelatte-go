package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

type UserRepository struct {
	db              *sqlx.DB
	addressRepo     ports.IAddressRepository
	paymentCardRepo ports.IPaymentCardRepository
}

func (ur *UserRepository) SelectAllUsers() ([]entities.User, error) {
	var temporaryUsers []models.TemporaryUserModel

	query := `select 
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
    		on u.ID = up.UserID where u.Status=true and ua.Status=true and up.Status=true order by u.ID`
	err := ur.db.Select(&temporaryUsers, query)
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
	var userModel models.UserModel

	err := ur.db.Get(&userModel, "select * from user u where u.ID=? and u.Status=true", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WrapError(errors.ErrRecordNotFound, err.Error())
		}
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	user := mappers.FromUserModelToUser(userModel)

	addresses, err := ur.addressRepo.SelectAddressesByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	user.Addresses = addresses

	cards, err := ur.paymentCardRepo.SelectCardsByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	user.PaymentCards = cards

	return user, nil
}

func (ur *UserRepository) SelectUserByEmail(email string) (*entities.User, error) {
	var userModel models.UserModel

	err := ur.db.Get(&userModel, "select * from user u where u.Email=? and u.Status=true", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WrapError(errors.ErrRecordNotFound, err.Error())
		}
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	user := mappers.FromUserModelToUser(userModel)

	addresses, err := ur.addressRepo.SelectAddressesByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	user.Addresses = addresses

	cards, err := ur.paymentCardRepo.SelectCardsByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	user.PaymentCards = cards

	return user, nil
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

func NewUserRepository(db *sqlx.DB,
	addressRepo ports.IAddressRepository,
	paymentCardRepo ports.IPaymentCardRepository) *UserRepository {

	return &UserRepository{
		db:              db,
		addressRepo:     addressRepo,
		paymentCardRepo: paymentCardRepo,
	}
}

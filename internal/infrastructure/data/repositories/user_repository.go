package repositories

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	core "github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

type UserRepository struct {
	db *sqlx.DB
}

var (
	insertUserError = errors.New("errors in inserting new user")
	selectUserError = errors.New("errors in selecting user(s)")
	updateUserError = errors.New("errors in updating user")
)

const selectTemporaryUsers = `select u.ID               as 'UserID',
					   u.Username         as 'UserUsername',
					   u.Name             as 'UserName',
					   u.Surname          as 'UserSurname',
					   u.PhoneNumber      as 'UserPhoneNumber',
					   u.Email            as 'UserEmail',
					   u.Password         as 'UserPassword',
					   u.RoleCode         as 'UserRoleCode',
					   u.Status           as 'UserStatus',
					   u.CreatedAt        as 'UserCreatedAt',
					   u.UpdatedAt        as 'UserUpdatedAt',
					   ua.ID              as 'AddressID',
					   ua.Type            as 'AddressType',
					   ua.ProvinceID      as 'AddressProvinceID',
					   ua.CityID          as 'AddressCityID',
					   ua.PostalCode      as 'AddressPostalCode',
					   ua.Detail          as 'AddressDetail',
					   ua.Status          as 'AddressStatus',
					   ua.CreatedAt       as 'AddressCreatedAt',
					   ua.UpdatedAt       as 'AddressUpdatedAt',
					   up.ID              as 'CardID',
					   up.Type            as 'CardType',
					   up.Company         as 'CardCompany',
					   up.HolderName      as 'CardHolderName',
					   up.Number          as 'CardNumber',
					   up.ExpirationYear  as 'CardExpirationYear',
					   up.ExpirationMonth as 'CardExpirationMonth',
					   up.CVV             as 'CardCVV',
					   up.Status          as 'CardStatus',
					   up.CreatedAt       as 'CardCreatedAt',
					   up.UpdatedAt       as 'CardUpdatedAt'
				from user u
						 left join useraddress ua on u.ID = ua.UserID
						 left join userpaymentcard up on u.ID = up.UserID
				where u.Status = true
				  and (ua.Status = true or ua.Status is null)
				  and (up.Status = true or up.Status is null)`

func (r *UserRepository) SelectAllUsers() ([]entities.User, error) {
	users := make([]entities.User, 0)

	var temporaryUsers []models.TemporaryUserModel

	err := r.db.Select(&temporaryUsers, selectTemporaryUsers)
	if err != nil {
		return nil, core.NewAppError(
			errors.Join(selectUserError, err),
			core.RepositoryError,
		)
	}

	if temporaryUsers == nil {
		return nil, core.NewAppErrorWithType(core.NotFoundError)
	}

	users = mappers.FromTemporaryUsersModelToUserSlice(temporaryUsers)
	return users, nil
}

func (r *UserRepository) SelectUserByID(userID int) (*entities.User, error) {
	var temporaryUsers []models.TemporaryUserModel

	err := r.db.Select(
		&temporaryUsers,
		selectTemporaryUsers+" and u.ID=?",
		userID,
	)
	if err != nil {
		return nil, core.NewAppError(
			errors.Join(selectUserError, err),
			core.RepositoryError,
		)
	}

	if temporaryUsers == nil {
		return nil, core.NewAppErrorWithType(core.NotFoundError)
	}

	users := mappers.FromTemporaryUsersModelToUserSlice(temporaryUsers)
	return &users[0], nil
}

func (r *UserRepository) SelectUserByEmail(email string) (
	*entities.User,
	error,
) {
	var temporaryUsers []models.TemporaryUserModel

	err := r.db.Select(
		&temporaryUsers,
		selectTemporaryUsers+" and u.Email=?",
		email,
	)
	if err != nil {
		return nil, core.NewAppError(
			errors.Join(selectUserError, err),
			core.RepositoryError,
		)
	}

	if temporaryUsers == nil {
		return nil, core.NewAppErrorWithType(core.NotFoundError)
	}

	users := mappers.FromTemporaryUsersModelToUserSlice(temporaryUsers)
	return &users[0], nil
}

func (r *UserRepository) InsertUser(user entities.User) (
	*entities.User,
	error,
) {
	userModel := mappers.FromUserToUserModel(user)

	result, err := r.db.Exec(
		`insert into user (
                  Username, 
                  Name, 
                  Surname, 
                  PhoneNumber, 
                  Email, 
                  Password, 
                  RoleCode
        	) values (?,?,?,?,?,?,?)`,
		userModel.Username,
		userModel.Name,
		userModel.Surname,
		userModel.PhoneNumber,
		userModel.Email,
		userModel.Password,
		userModel.RoleCode,
	)
	if err != nil {
		return nil, core.NewAppError(
			errors.Join(insertUserError, err),
			core.RepositoryError,
		)
	}

	lastUserID, _ := result.LastInsertId()

	userModel.ID = int(lastUserID)
	return mappers.FromUserModelToUser(*userModel), nil
}

func (r *UserRepository) UpdateUser(userID int, user entities.User) error {
	userModel := mappers.FromUserToUserModel(user)

	query := `update user set 
                Username=?, 
                Name=?, 
                Surname=?, 
                PhoneNumber=? 
            where ID=?`

	_, err := r.db.Exec(
		query,
		userModel.Name,
		userModel.Surname,
		userModel.PhoneNumber,
		userID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.NewAppErrorWithType(core.NotFoundError)
		}
		return core.NewAppError(
			errors.Join(updateUserError, err),
			core.RepositoryError,
		)
	}

	return nil
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

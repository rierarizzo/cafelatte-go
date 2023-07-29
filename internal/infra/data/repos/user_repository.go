package repos

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infra/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infra/data/models"
)

// UserRepo represents a repository for user-related operations.
type UserRepo struct {
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

// SelectUsers retrieves a list of users from the database and returns the
// list of users if successful, along with any error encountered during the
// process.
func (r *UserRepo) SelectUsers() ([]entities.User, error) {
	users := make([]entities.User, 0)

	var temporaryUsers []models.TemporaryUserModel

	err := r.db.Select(&temporaryUsers, selectTemporaryUsers)
	if err != nil {
		return nil, domain.NewAppError(
			errors.Join(selectUserError, err),
			domain.RepositoryError,
		)
	}

	if temporaryUsers == nil {
		return nil, domain.NewAppErrorWithType(domain.NotFoundError)
	}

	users = mappers.FromTemporaryUsersModelToUserSlice(temporaryUsers)
	return users, nil
}

// SelectUserByID retrieves a user from the database based on the provided
// user ID and returns the user if found, along with any error encountered
// during the process.
func (r *UserRepo) SelectUserByID(userID int) (*entities.User, error) {
	var temporaryUsers []models.TemporaryUserModel

	err := r.db.Select(
		&temporaryUsers,
		selectTemporaryUsers+" and u.ID=?",
		userID,
	)
	if err != nil {
		return nil, domain.NewAppError(
			errors.Join(selectUserError, err),
			domain.RepositoryError,
		)
	}

	if temporaryUsers == nil {
		return nil, domain.NewAppErrorWithType(domain.NotFoundError)
	}

	users := mappers.FromTemporaryUsersModelToUserSlice(temporaryUsers)
	return &users[0], nil
}

// SelectUserByEmail retrieves a user from the database based on the
// provided email and returns the user if found, along with any error
// encountered during the process.
func (r *UserRepo) SelectUserByEmail(email string) (
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
		return nil, domain.NewAppError(
			errors.Join(selectUserError, err),
			domain.RepositoryError,
		)
	}

	if temporaryUsers == nil {
		return nil, domain.NewAppErrorWithType(domain.NotFoundError)
	}

	users := mappers.FromTemporaryUsersModelToUserSlice(temporaryUsers)
	return &users[0], nil
}

// InsertUser inserts a new user into the database and returns the inserted
// user if successful, along with any error encountered during the process.
func (r *UserRepo) InsertUser(user entities.User) (
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
		return nil, domain.NewAppError(
			errors.Join(insertUserError, err),
			domain.RepositoryError,
		)
	}

	lastUserID, _ := result.LastInsertId()

	userModel.ID = int(lastUserID)
	return mappers.FromUserModelToUser(*userModel), nil
}

// UpdateUser updates the details of a user in the database based on the
// provided user ID and user object and returns an error, if any,
// encountered during the process.
func (r *UserRepo) UpdateUser(userID int, user entities.User) error {
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
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}
		return domain.NewAppError(
			errors.Join(updateUserError, err),
			domain.RepositoryError,
		)
	}

	return nil
}

func NewUserRepository(db *sqlx.DB) *UserRepo {
	return &UserRepo{db}
}

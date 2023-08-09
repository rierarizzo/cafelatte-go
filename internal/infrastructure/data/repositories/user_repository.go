package repositories

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/constants"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	db *sqlx.DB
}

// SelectUsers retrieves a list of users from the database and returns the
// list of users if successful, along with any error encountered during the
// process.
func (r *UserRepository) SelectUsers() ([]entities.User, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	var usersModel []models.UserModel
	var query = "select * from user where Status=true"

	err := r.db.Select(&usersModel, query)
	if err != nil {
		log.Error(err)
		return nil, domain.NewAppError(selectUserError, domain.RepositoryError)
	}

	if usersModel == nil {
		return nil, domain.NewAppErrorWithType(domain.NotFoundError)
	}

	return mappers.FromUserModelSliceToUserSlice(usersModel), nil
}

// SelectUserByID retrieves a user from the database based on the provided
// user ID and returns the user if found, along with any error encountered
// during the process.
func (r *UserRepository) SelectUserByID(userID int) (*entities.User, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	var userModel models.UserModel
	var query = "select * from user u where u.Status=true and u.ID=?"

	err := r.db.Get(&userModel, query, userID)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return nil, domain.NewAppError(selectUserError, domain.RepositoryError)
	}

	user := mappers.FromUserModelToUser(userModel)
	return &user, nil
}

// SelectUserByEmail retrieves a user from the database based on the
// provided email and returns the user if found, along with any error
// encountered during the process.
func (r *UserRepository) SelectUserByEmail(email string) (*entities.User, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	var userModel models.UserModel
	var query = "select * from user u where u.Status=true and u.Email=?"

	err := r.db.Get(&userModel, query, email)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return nil, domain.NewAppError(selectUserError, domain.RepositoryError)
	}

	user := mappers.FromUserModelToUser(userModel)
	return &user, nil
}

// InsertUser inserts a new user into the database and returns the inserted
// user if successful, along with any error encountered during the process.
func (r *UserRepository) InsertUser(user entities.User) (*entities.User, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	var userModel = mappers.FromUserToUserModel(user)
	var query = `insert into user (
                  Username, 
                  Name, 
                  Surname, 
                  PhoneNumber, 
                  Email, 
                  Password, 
                  RoleCode) 
			values (?,?,?,?,?,?,?)`

	result, err := r.db.Exec(query, userModel.Username, userModel.Name,
		userModel.Surname, userModel.PhoneNumber, userModel.Email,
		userModel.Password, userModel.RoleCode)
	if err != nil {
		log.Error(err)
		return nil, domain.NewAppError(insertUserError, domain.RepositoryError)
	}

	lastUserID, _ := result.LastInsertId()
	userModel.ID = int(lastUserID)

	u := mappers.FromUserModelToUser(userModel)
	userToReturn := &u
	return userToReturn, nil
}

// UpdateUser updates the details of a user in the database based on the
// provided user ID and user object and returns an error, if any,
// encountered during the process.
func (r *UserRepository) UpdateUser(userID int,
	user entities.User) *domain.AppError {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	var userModel = mappers.FromUserToUserModel(user)
	var query = `update user set 
                Username=?, 
                Name=?, 
                Surname=?, 
                PhoneNumber=? 
            where ID=?`

	_, err := r.db.Exec(query, userModel.Username, userModel.Name,
		userModel.Surname, userModel.PhoneNumber, userID)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return domain.NewAppError(updateUserError, domain.RepositoryError)
	}

	return nil
}

func (r *UserRepository) DeleteUser(userID int) *domain.AppError {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	var query = `update user set Status=false where ID=?`

	_, err := r.db.Exec(query, userID)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return domain.NewAppError(deleteUserError, domain.RepositoryError)
	}

	return nil
}

var (
	insertUserError = errors.New("insert user error")
	selectUserError = errors.New("select user error")
	updateUserError = errors.New("update user error")
	deleteUserError = errors.New("delete user error")
)

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

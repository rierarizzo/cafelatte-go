package user

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
)

var (
	insertUserError = errors.New("insert usermanager error")
	selectUserError = errors.New("select usermanager error")
	updateUserError = errors.New("update usermanager error")
	deleteUserError = errors.New("delete usermanager error")
)

type Repository struct {
	db *sqlx.DB
}

// SelectUsers retrieves a list of users from the database and returns the
// list of users if successful, along with any error encountered during the
// process.
func (repository *Repository) SelectUsers() ([]domain.User, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var usersModel []Model
	var query = "select * from User where Status=true"

	err := repository.db.Select(&usersModel, query)
	if err != nil {
		log.Error(err)
		return nil, domain.NewAppError(selectUserError, domain.RepositoryError)
	}

	if usersModel == nil {
		return nil, domain.NewAppErrorWithType(domain.NotFoundError)
	}

	return fromModelsToUsers(usersModel), nil
}

// SelectUserById retrieves a usermanager from the database based on the provided
// usermanager ID and returns the usermanager if found, along with any error encountered
// during the process.
func (repository *Repository) SelectUserById(userID int) (*domain.User, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var userModel Model
	var query = "select * from User u where u.Status=true and u.ID=?"

	err := repository.db.Get(&userModel, query, userID)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return nil, domain.NewAppError(selectUserError, domain.RepositoryError)
	}

	user := fromModelToUser(userModel)
	return &user, nil
}

// SelectUserByEmail retrieves a usermanager from the database based on the
// provided email and returns the usermanager if found, along with any error
// encountered during the process.
func (repository *Repository) SelectUserByEmail(email string) (*domain.User, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var userModel Model
	var query = "select * from User u where u.Status=true and u.Email=?"

	err := repository.db.Get(&userModel, query, email)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return nil, domain.NewAppError(selectUserError, domain.RepositoryError)
	}

	user := fromModelToUser(userModel)
	return &user, nil
}

// InsertUser inserts a new usermanager into the database and returns the inserted
// usermanager if successful, along with any error encountered during the process.
func (repository *Repository) InsertUser(user domain.User) (*domain.User, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var userModel = fromUserToModel(user)
	var query = `insert into User (
                  Username, 
                  Name, 
                  Surname, 
                  PhoneNumber, 
                  Email, 
                  Password, 
                  RoleCode) 
			values (?,?,?,?,?,?,?)`

	result, err := repository.db.Exec(query, userModel.Username, userModel.Name,
		userModel.Surname, userModel.PhoneNumber, userModel.Email,
		userModel.Password, userModel.RoleCode)
	if err != nil {
		log.Error(err)
		return nil, domain.NewAppError(insertUserError, domain.RepositoryError)
	}

	lastUserID, _ := result.LastInsertId()
	userModel.ID = int(lastUserID)

	u := fromModelToUser(userModel)
	userToReturn := &u
	return userToReturn, nil
}

// UpdateUserById updates the details of a usermanager in the database based on the
// provided usermanager ID and usermanager object and returns an error, if any,
// encountered during the process.
func (repository *Repository) UpdateUserById(userID int,
	user domain.User) *domain.AppError {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var userModel = fromUserToModel(user)
	var query = `update User set 
                Username=?, 
                Name=?, 
                Surname=?, 
                PhoneNumber=? 
            where ID=?`

	_, err := repository.db.Exec(query, userModel.Username, userModel.Name,
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

func (repository *Repository) DeleteUserById(userID int) *domain.AppError {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var query = `update User set Status=false where ID=?`

	_, err := repository.db.Exec(query, userID)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return domain.NewAppError(deleteUserError, domain.RepositoryError)
	}

	return nil
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}

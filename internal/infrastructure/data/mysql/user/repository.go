package user

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	userDomain "github.com/rierarizzo/cafelatte/internal/domain/user"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	db *sqlx.DB
}

// SelectUsers retrieves a list of users from the database and returns the
// list of users if successful, along with any error encountered during the
// process.
func (r *Repository) SelectUsers() ([]userDomain.User, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var usersModel []Model
	var query = "select * from User where Status=true"

	err := r.db.Select(&usersModel, query)
	if err != nil {
		log.Error(err)
		return nil, domain.NewAppError(selectUserError, domain.RepositoryError)
	}

	if usersModel == nil {
		return nil, domain.NewAppErrorWithType(domain.NotFoundError)
	}

	return fromModelsToUsers(usersModel), nil
}

// SelectUserByID retrieves a user from the database based on the provided
// user ID and returns the user if found, along with any error encountered
// during the process.
func (r *Repository) SelectUserByID(userID int) (*userDomain.User, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var userModel Model
	var query = "select * from User u where u.Status=true and u.ID=?"

	err := r.db.Get(&userModel, query, userID)
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

// SelectUserByEmail retrieves a user from the database based on the
// provided email and returns the user if found, along with any error
// encountered during the process.
func (r *Repository) SelectUserByEmail(email string) (*userDomain.User, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var userModel Model
	var query = "select * from User u where u.Status=true and u.Email=?"

	err := r.db.Get(&userModel, query, email)
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

// InsertUser inserts a new user into the database and returns the inserted
// user if successful, along with any error encountered during the process.
func (r *Repository) InsertUser(user userDomain.User) (*userDomain.User, *domain.AppError) {
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

	result, err := r.db.Exec(query, userModel.Username, userModel.Name,
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

// UpdateUser updates the details of a user in the database based on the
// provided user ID and user object and returns an error, if any,
// encountered during the process.
func (r *Repository) UpdateUser(userID int,
	user userDomain.User) *domain.AppError {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var userModel = fromUserToModel(user)
	var query = `update User set 
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

func (r *Repository) DeleteUser(userID int) *domain.AppError {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var query = `update User set Status=false where ID=?`

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

func NewUserRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}

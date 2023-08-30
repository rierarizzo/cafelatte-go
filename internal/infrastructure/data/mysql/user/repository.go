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
	insertUserError = errors.New("insert user error")
	selectUserError = errors.New("select user error")
	updateUserError = errors.New("update user error")
	deleteUserError = errors.New("delete user error")
)

type Repository struct {
	db *sqlx.DB
}

// SelectUsers retrieves a list of users from the database and returns the
// list of users if successful, along with any error encountered during the
// process.
func (r *Repository) SelectUsers() ([]domain.User, *domain.AppError) {
	log := logrus.WithField(misc.RequestIdKey, request.Id())

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

// SelectUserById retrieves a usermanager from the database based on the provided
// usermanager Id and returns the usermanager if found, along with any error encountered
// during the process.
func (r *Repository) SelectUserById(userId int) (*domain.User,
	*domain.AppError) {
	log := logrus.WithField(misc.RequestIdKey, request.Id())

	var userModel Model
	var query = "select * from User u where u.Status=true and u.Id=?"

	err := r.db.Get(&userModel, query, userId)
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
func (r *Repository) SelectUserByEmail(email string) (*domain.User,
	*domain.AppError) {
	log := logrus.WithField(misc.RequestIdKey, request.Id())

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

// InsertUser inserts a new usermanager into the database and returns the inserted
// usermanager if successful, along with any error encountered during the process.
func (r *Repository) InsertUser(user domain.User) (*domain.User,
	*domain.AppError) {
	log := logrus.WithField(misc.RequestIdKey, request.Id())

	var userModel = fromUserToModel(user)
	var query = `insert into User (Username, Name, Surname, PhoneNumber, Email, 
        Password, RoleCode) values (?,?,?,?,?,?,?)`

	result, err := r.db.Exec(query, userModel.Username, userModel.Name,
		userModel.Surname, userModel.PhoneNumber, userModel.Email,
		userModel.Password, userModel.RoleCode)
	if err != nil {
		log.Error(err)
		return nil, domain.NewAppError(insertUserError, domain.RepositoryError)
	}

	lastUserId, _ := result.LastInsertId()
	userModel.Id = int(lastUserId)

	u := fromModelToUser(userModel)
	userToReturn := &u
	return userToReturn, nil
}

// UpdateUserById updates the details of a usermanager in the database based on the
// provided usermanager Id and usermanager object and returns an error, if any,
// encountered during the process.
func (r *Repository) UpdateUserById(userId int,
	user domain.User) *domain.AppError {
	log := logrus.WithField(misc.RequestIdKey, request.Id())

	var userModel = fromUserToModel(user)
	var query = `update User set Username=?, Name=?, Surname=?, PhoneNumber=? where Id=?`
	_, err := r.db.Exec(query, userModel.Username, userModel.Name,
		userModel.Surname, userModel.PhoneNumber, userId)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return domain.NewAppError(updateUserError, domain.RepositoryError)
	}

	return nil
}

func (r *Repository) DeleteUserById(userId int) *domain.AppError {
	log := logrus.WithField(misc.RequestIdKey, request.Id())

	var query = `update User set Status=false where Id=?`

	_, err := r.db.Exec(query, userId)
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

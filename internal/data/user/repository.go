package user

import (
	"database/sql"
	"errors"
	sqlUtil "github.com/rierarizzo/cafelatte/pkg/utils/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
)

const NotFoundMsg = "user not found"

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) SelectUsers() ([]domain.User, *domain.AppError) {
	var usersModel []Model

	var query = `
		SELECT * FROM User WHERE Status=TRUE
	`
	err := r.db.Select(&usersModel, query)
	if err != nil {
		appErr := domain.NewAppError(err, domain.RepositoryError)
		return nil, appErr
	}

	if usersModel == nil {
		return []domain.User{}, nil
	}

	users := fromModelsToUsers(usersModel)
	return users, nil
}

func (r *Repository) SelectUserById(userId int) (*domain.User, *domain.AppError) {
	var userModel Model

	var query = `
		SELECT * FROM User u WHERE u.Id=? AND u.Status=TRUE
	`
	err := r.db.Get(&userModel, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			appErr := domain.NewAppError(NotFoundMsg, domain.NotFoundError)
			return nil, appErr
		}

		appErr := domain.NewAppError(err, domain.RepositoryError)
		return nil, appErr
	}

	user := fromModelToUser(userModel)
	return &user, nil
}

func (r *Repository) SelectUserByEmail(email string) (*domain.User, *domain.AppError) {
	var userModel Model

	var query = `
		SELECT * FROM User u WHERE u.Email=? AND u.Status=TRUE
	`
	err := r.db.Get(&userModel, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			appErr := domain.NewAppError(NotFoundMsg, domain.NotFoundError)
			return nil, appErr
		}

		appErr := domain.NewAppError(err, domain.RepositoryError)
		return nil, appErr
	}

	user := fromModelToUser(userModel)
	return &user, nil
}

func (r *Repository) InsertUser(user domain.User) (*domain.User, *domain.AppError) {
	var model = fromUserToModel(user)

	var query = `
		INSERT INTO User (Username, Name, Surname, PhoneNumber, Email, Password, RoleCode) 
		VALUES (?,?,?,?,?,?,?)
	`
	result, err := r.db.Exec(query, model.Username, model.Name, model.Surname, model.PhoneNumber, model.Email,
		model.Password, model.RoleCode)
	if err != nil {
		appErr := domain.NewAppError(err, domain.RepositoryError)
		return nil, appErr
	}

	userId, appErr := sqlUtil.GetLastInsertedId(result)
	if appErr != nil {
		return nil, appErr
	}
	model.Id = userId

	u := fromModelToUser(model)
	userToReturn := &u
	return userToReturn, nil
}

func (r *Repository) UpdateUserById(userId int, user domain.User) *domain.AppError {
	var model = fromUserToModel(user)

	var query = `
		UPDATE User SET Username=?, Name=?, Surname=?, PhoneNumber=? WHERE Id=?
	`
	_, err := r.db.Exec(query, model.Username, model.Name, model.Surname, model.PhoneNumber, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			appErr := domain.NewAppError(NotFoundMsg, domain.NotFoundError)
			return appErr
		}

		appErr := domain.NewAppError(err, domain.RepositoryError)
		return appErr
	}

	return nil
}

func (r *Repository) DeleteUserById(userId int) *domain.AppError {
	var query = `
		UPDATE User SET Status=FALSE WHERE Id=?
	`
	_, err := r.db.Exec(query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			appErr := domain.NewAppError(NotFoundMsg, domain.NotFoundError)
			return appErr
		}

		appErr := domain.NewAppError(err, domain.RepositoryError)
		return appErr
	}

	return nil
}

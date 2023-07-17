package repositories

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
	"github.com/sirupsen/logrus"
)

const (
	getUserErrorMsg    = "error while retrieving user from database: %v"
	getUsersErrorMsg   = "error while retrieving users from database: %v"
	insertUserErrorMsg = "error while inserting user to database: %v"
	updateUserErrorMsg = "error while updating user in database: %v"
)

type UserRepository struct {
	db *sqlx.DB
}

func (ur *UserRepository) GetAllUsers() ([]entities.User, error) {
	var userModel []models.UserModel

	query := "SELECT * FROM user"
	err := ur.db.Select(&userModel, query)
	if err != nil {
		logrus.Errorf(getUsersErrorMsg, err)
		if err == sql.ErrNoRows {
			var emptyUsers []entities.User
			return emptyUsers, nil
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
	userModel := models.UserModel{}

	query := "SELECT * FROM user u WHERE u.id=?"
	err := ur.db.Get(&userModel, query, userID)
	if err != nil {
		logrus.Errorf(getUserErrorMsg, err)
		return nil, handleSQLError(err)
	}

	return userModel.ToUserCore(), nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	userModel := models.UserModel{}

	query := "SELECT * FROM user u WHERE u.email=?"
	err := ur.db.Get(&userModel, query, email)
	if err != nil {
		logrus.Errorf(getUserErrorMsg, err)
		return nil, handleSQLError(err)
	}

	return userModel.ToUserCore(), nil
}

func (ur *UserRepository) CreateUser(user entities.User) (*entities.User, error) {
	var userModel models.UserModel
	userModel.LoadFromUserCore(user)

	query := "INSERT INTO user (name, surname, phone_number, email, password) VALUES (?,?,?,?,?)"

	result, err := ur.db.Exec(query,
		userModel.Name, userModel.Surname, userModel.PhoneNumber, userModel.Email, userModel.Password)
	if err != nil {
		logrus.Errorf(insertUserErrorMsg, err)
		return nil, handleSQLError(err)
	}

	lastID, _ := result.LastInsertId()
	userModel.ID = int(lastID)

	return userModel.ToUserCore(), nil
}

func (ur *UserRepository) UpdateUser(userID int, user entities.User) error {
	var userModel models.UserModel
	userModel.LoadFromUserCore(user)

	query := "UPDATE user SET name=?, surname=?, phone_number=?, email=?, password=? WHERE id=?"

	_, err := ur.db.Exec(query, user.Name, user.Surname, user.PhoneNumber, user.Email, user.Password, userID)
	if err != nil {
		logrus.Errorf(updateUserErrorMsg, err)
		return handleSQLError(err)
	}

	return nil
}

func handleSQLError(sqlError error) error {
	switch sqlError {
	case sql.ErrNoRows:
		return core.ErrRecordNotFound
	default:
		return core.ErrUnexpected
	}
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

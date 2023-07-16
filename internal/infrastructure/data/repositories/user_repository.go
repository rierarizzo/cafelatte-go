package repositories

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	db *sqlx.DB
}

const (
	getUserErrorMsg    = "error while retrieving user from database: %v"
	insertUserErrorMsg = "error while inserting user to database: %v"
)

func (ur *UserRepository) GetUserById(id int) (*entities.User, error) {
	userModel := models.UserModel{}

	query := "SELECT * FROM user u WHERE u.id = ?"
	err := ur.db.Get(&userModel, query, id)
	if err != nil {
		logrus.Errorf(getUserErrorMsg, err)
		return nil, handleSQLError(err)
	}

	return userModel.ToUserCore(), nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	userModel := models.UserModel{}

	query := "SELECT * FROM user u WHERE u.email = ?"
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

	query := `INSERT INTO user (name, surname, phone_number, email, password)
				VALUES (?, ?, ?, ?, ?)`

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

func handleSQLError(sqlError error) error {
	switch sqlError {
	case sql.ErrNoRows:
		return core.RecordNotFound
	default:
		return core.Unexpected
	}
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

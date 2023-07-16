package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func (ur *UserRepository) GetUserById(id int) (*entities.User, error) {
	userModel := models.UserModel{}

	query := "SELECT * FROM user u WHERE u.id = ?"
	err := ur.db.Get(&userModel, query, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return userModel.ToUserCore(), nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	userModel := models.UserModel{}

	query := "SELECT * FROM user u WHERE u.email = ?"
	err := ur.db.Get(&userModel, query, email)
	if err != nil {
		log.Println(err)
		return nil, err
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
		log.Println(err)
		return nil, err
	}

	lastID, _ := result.LastInsertId()
	userModel.ID = int(lastID)

	return userModel.ToUserCore(), nil
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
	"sync"
)

type UserRepository struct {
	db              *sqlx.DB
	addressRepo     ports.IAddressRepository
	paymentCardRepo ports.IPaymentCardRepository
}

func (ur *UserRepository) SelectAllUsers() ([]entities.User, error) {
	var usersModel []models.UserModel

	query := "select * from User u where u.Status=true"
	err := ur.db.Select(&usersModel, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []entities.User{}, nil
		} else {
			return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
		}
	}

	var users []entities.User
	for _, k := range usersModel {
		users = append(users, *mappers.FromUserModelToUser(k))
	}

	sem := make(chan struct{}, 3)

	errCh := make(chan error, len(users))
	var wg sync.WaitGroup

	for i, v := range users {
		wg.Add(1)
		sem <- struct{}{}

		go func(userIndex int, user entities.User) {
			defer func() {
				wg.Done()
				<-sem
			}()

			addresses, err := ur.addressRepo.SelectAddressesByUserID(user.ID)
			if err != nil {
				errCh <- err
				return
			}

			cards, err := ur.paymentCardRepo.SelectCardsByUserID(user.ID)
			if err != nil {
				errCh <- err
				return
			}

			users[userIndex].Addresses = addresses
			users[userIndex].PaymentCards = cards
		}(i, v)

	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	return users, nil
}

func (ur *UserRepository) SelectUserByID(userID int) (*entities.User, error) {
	var userModel models.UserModel

	err := ur.db.Get(&userModel, "select * from User u where u.ID=? and u.Status=true", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WrapError(errors.ErrRecordNotFound, err.Error())
		}
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	user := mappers.FromUserModelToUser(userModel)

	addresses, err := ur.addressRepo.SelectAddressesByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	user.Addresses = addresses

	cards, err := ur.paymentCardRepo.SelectCardsByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	user.PaymentCards = cards

	return user, nil
}

func (ur *UserRepository) SelectUserByEmail(email string) (*entities.User, error) {
	var userModel models.UserModel

	err := ur.db.Get(&userModel, "select * from User u where u.Email=? and u.Status=true", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WrapError(errors.ErrRecordNotFound, err.Error())
		}
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	user := mappers.FromUserModelToUser(userModel)

	addresses, err := ur.addressRepo.SelectAddressesByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	user.Addresses = addresses

	cards, err := ur.paymentCardRepo.SelectCardsByUserID(user.ID)
	if err != nil {
		return nil, err
	}
	user.PaymentCards = cards

	return user, nil
}

func (ur *UserRepository) InsertUser(user entities.User) (*entities.User, error) {
	userModel := mappers.FromUserToUserModel(user)

	result, err := ur.db.Exec(
		`insert into User (Username, Name, Surname, PhoneNumber, Email, Password, RoleCode) 
			values (?,?,?,?,?,?,?)`,
		userModel.Username, userModel.Name, userModel.Surname, userModel.PhoneNumber,
		userModel.Email, userModel.Password, userModel.RoleCode)
	if err != nil {
		return nil, errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	lastUserID, _ := result.LastInsertId()

	userModel.ID = int(lastUserID)
	return mappers.FromUserModelToUser(*userModel), nil
}

func (ur *UserRepository) UpdateUser(userID int, user entities.User) error {
	userModel := mappers.FromUserToUserModel(user)

	query := "update User set Username=?, Name=?, Surname=?, PhoneNumber=? where ID=?"

	_, err := ur.db.Exec(query, userModel.Name, userModel.Surname, userModel.PhoneNumber, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.WrapError(errors.ErrRecordNotFound, err.Error())
		}
		return errors.WrapError(errors.ErrUnexpected, err.Error())
	}

	return nil
}

func NewUserRepository(db *sqlx.DB,
	addressRepo ports.IAddressRepository,
	paymentCardRepo ports.IPaymentCardRepository) *UserRepository {

	return &UserRepository{
		db:              db,
		addressRepo:     addressRepo,
		paymentCardRepo: paymentCardRepo,
	}
}

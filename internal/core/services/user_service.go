package services

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	userRepo ports.IUserRepository
}

func (us *UserService) SignUp(user entities.User) (*entities.AuthorizedUser, error) {
	if !user.IsValidUser() {
		logrus.Error(errors.ErrInvalidUserData.Error())
		return nil, errors.ErrInvalidUserData
	}

	hashedPassword, err := utils.HashText(user.Password)
	if err != nil {
		return nil, errors.ErrUnexpected
	}
	user.SetPassword(hashedPassword)

	retrievedUser, err := us.userRepo.InsertUser(user)
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	token, err := utils.CreateJWTToken(*retrievedUser)
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	authorizedUser := entities.AuthorizedUser{
		User:        *retrievedUser,
		AccessToken: *token,
	}

	return &authorizedUser, nil
}

func (us *UserService) SignIn(email, password string) (*entities.AuthorizedUser, error) {
	retrievedUser, err := us.userRepo.SelectUserByEmail(email)
	if err != nil {
		return nil, errors.ErrUnauthorizedUser
	}

	if !utils.CheckTextHash(retrievedUser.Password, password) {
		return nil, errors.ErrUnauthorizedUser
	}

	token, err := utils.CreateJWTToken(*retrievedUser)
	if err != nil {
		return nil, errors.ErrUnauthorizedUser
	}

	authorizedUser := entities.AuthorizedUser{
		User:        *retrievedUser,
		AccessToken: *token,
	}

	return &authorizedUser, nil
}

func (us *UserService) GetAllUsers() ([]entities.User, error) {
	return us.userRepo.SelectAllUsers()
}

func (us *UserService) FindUserByID(id int) (*entities.User, error) {
	return us.userRepo.SelectUserByID(id)
}

func (us *UserService) UpdateUser(userID int, user entities.User) error {
	return us.userRepo.UpdateUser(userID, user)
}

func (us *UserService) AddUserAddresses(userID int, addresses []entities.Address) ([]entities.Address, error) {
	return us.userRepo.InsertUserAddresses(userID, addresses)
}

func (us *UserService) AddUserPaymentCard(userID int, cards []entities.PaymentCard) ([]entities.PaymentCard, error) {
	return us.userRepo.InsertUserPaymentCards(userID, cards)
}

func NewUserService(userRepo ports.IUserRepository) *UserService {
	return &UserService{userRepo}
}

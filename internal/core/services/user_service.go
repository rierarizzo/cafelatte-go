package services

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type UserService struct {
	userRepo ports.IUserRepository
}

func (us *UserService) SignUp(user entities.User) (*entities.AuthorizedUser, error) {
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashText(user.Password)
	if err != nil {
		return nil, err
	}
	user.SetPassword(hashedPassword)

	retrievedUser, err := us.userRepo.InsertUser(user)
	if err != nil {
		return nil, err
	}

	token, err := utils.CreateJWTToken(*retrievedUser)
	if err != nil {
		return nil, err
	}

	authorizedUser := entities.AuthorizedUser{
		User:        *retrievedUser,
		AccessToken: *token,
	}

	return &authorizedUser, nil
}

func (us *UserService) SignIn(email, password string) (*entities.AuthorizedUser, error) {
	const incorrectEmailOrPasswordMsg = "incorrect email or password"

	retrievedUser, err := us.userRepo.SelectUserByEmail(email)
	if err != nil {
		if utils.CompareErrors(err, errors.ErrRecordNotFound) {
			return nil, errors.WrapError(errors.ErrUnauthorizedUser, incorrectEmailOrPasswordMsg)
		}

		return nil, err
	}

	if !utils.CheckTextHash(retrievedUser.Password, password) {
		return nil, errors.WrapError(errors.ErrUnauthorizedUser, incorrectEmailOrPasswordMsg)
	}

	token, err := utils.CreateJWTToken(*retrievedUser)
	if err != nil {
		return nil, err
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

func NewUserService(userRepo ports.IUserRepository) *UserService {
	return &UserService{userRepo}
}

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
	hashedPassword, err := utils.HashText(user.Password)
	if err != nil {
		return nil, errors.ErrUnexpected
	}
	user.SetPassword(hashedPassword)

	retrievedUser, err := us.userRepo.CreateUser(user)
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	token, err := utils.CreateJWTToken(*retrievedUser)
	if err != nil {
		return nil, errors.ErrUnexpected
	}

	authorizedUser := entities.AuthorizedUser{
		UserInfo:    *retrievedUser,
		AccessToken: *token,
	}

	return &authorizedUser, nil
}

func (us *UserService) SignIn(email, password string) (*entities.AuthorizedUser, error) {
	retrievedUser, err := us.userRepo.GetUserByEmail(email)
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
		UserInfo:    *retrievedUser,
		AccessToken: *token,
	}

	return &authorizedUser, nil
}

func (us *UserService) GetAllUsers() ([]entities.User, error) {
	return us.userRepo.GetAllUsers()
}

func (us *UserService) FindUserById(id int) (*entities.User, error) {
	return us.userRepo.GetUserById(id)
}

func (us *UserService) UpdateUser(userID int, user entities.User) error {
	return us.userRepo.UpdateUser(userID, user)
}

func NewUserService(userRepo ports.IUserRepository) *UserService {
	return &UserService{userRepo}
}

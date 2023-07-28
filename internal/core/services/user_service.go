package services

import (
	"fmt"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type UserService struct {
	userRepo ports.IUserRepository
}

func (s *UserService) SignUp(user entities.User) (*entities.AuthorizedUser, error) {
	if err := user.ValidateUser(); err != nil {
		return nil, errors.WrapError(
			err,
			fmt.Sprintf("user with username '%s' is not valid", user.Username),
		)
	}

	hash, err := utils.HashText(user.Password)
	if err != nil {
		return nil, errors.WrapError(err, "failed to hash password")
	}
	user.SetPassword(hash)

	retrUser, err := s.userRepo.InsertUser(user)
	if err != nil {
		return nil, errors.WrapError(
			err,
			fmt.Sprintf("failed to insert user with username '%s' into db", user.Username),
		)
	}

	token, err := utils.CreateJWTToken(*retrUser)
	if err != nil {
		return nil, errors.WrapError(err, "failed to create JWT")
	}

	return entities.NewAuthorizedUser(*retrUser, *token), nil
}

func (s *UserService) SignIn(email, password string) (*entities.AuthorizedUser, error) {
	const incorrectEmailOrPasswordMsg = "incorrect email or password"

	retrUser, err := s.userRepo.SelectUserByEmail(email)
	if err != nil {
		if errors.CompareErrors(err, errors.ErrRecordNotFound) {
			return nil, errors.WrapError(errors.ErrUnauthorizedUser, incorrectEmailOrPasswordMsg)
		}
		return nil, errors.WrapError(
			err,
			fmt.Sprintf("failed to select user with email '%s' from db", email),
		)
	}

	if !utils.CheckTextHash(retrUser.Password, password) {
		return nil, errors.WrapError(errors.ErrUnauthorizedUser, incorrectEmailOrPasswordMsg)
	}

	token, err := utils.CreateJWTToken(*retrUser)
	if err != nil {
		return nil, errors.WrapError(err, "failed to create JWT")
	}

	return entities.NewAuthorizedUser(*retrUser, *token), nil
}

func (s *UserService) GetAllUsers() ([]entities.User, error) {
	return s.userRepo.SelectAllUsers()
}

func (s *UserService) FindUserByID(id int) (*entities.User, error) {
	return s.userRepo.SelectUserByID(id)
}

func (s *UserService) UpdateUser(userID int, user entities.User) error {
	return s.userRepo.UpdateUser(userID, user)
}

func NewUserService(userRepo ports.IUserRepository) *UserService {
	return &UserService{userRepo}
}

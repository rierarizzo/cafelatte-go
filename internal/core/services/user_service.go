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

func (s *UserService) SignUp(user entities.User) (*entities.AuthorizedUser, error) {
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashText(user.Password)
	if err != nil {
		return nil, err
	}
	user.SetPassword(hashedPassword)

	retrievedUser, err := s.userRepo.InsertUser(user)
	if err != nil {
		return nil, err
	}

	token, err := utils.CreateJWTToken(*retrievedUser)
	if err != nil {
		return nil, err
	}

	return entities.NewAuthorizedUser(*retrievedUser, *token), nil
}

func (s *UserService) SignIn(email, password string) (*entities.AuthorizedUser, error) {
	const incorrectEmailOrPasswordMsg = "incorrect email or password"

	retrievedUser, err := s.userRepo.SelectUserByEmail(email)
	if err != nil {
		if errors.CompareErrors(err, errors.ErrRecordNotFound) {
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

	return entities.NewAuthorizedUser(*retrievedUser, *token), nil
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

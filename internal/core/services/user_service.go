package services

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	core "github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type UserService struct {
	userRepo ports.IUserRepository
}

func (s *UserService) SignUp(user entities.User) (
	*entities.AuthorizedUser,
	error,
) {
	if err := user.ValidateUser(); err != nil {
		return nil, core.NewAppError(err, core.ValidationError)
	}

	hash, err := utils.HashText(user.Password)
	if err != nil {
		return nil, core.NewAppErrorWithType(core.HashGenerationError)
	}
	user.Password = hash

	retrUser, err := s.userRepo.InsertUser(user)
	if err != nil {
		return nil, core.NewAppError(err, core.UnexpectedError)
	}

	token, err := utils.CreateJWTToken(*retrUser)
	if err != nil {
		return nil, core.NewAppError(err, core.TokenGenerationError)
	}

	return entities.NewAuthorizedUser(*retrUser, *token), nil
}

func (s *UserService) SignIn(email, password string) (
	*entities.AuthorizedUser,
	error,
) {
	retrUser, err := s.userRepo.SelectUserByEmail(email)
	if err != nil {
		var coreErr *core.AppError
		wrapped := errors.As(err, &coreErr)
		if wrapped && coreErr.Type == core.NotFoundError {
			return nil, core.NewAppErrorWithType(core.NotAuthorizedError)
		} else {
			return nil, core.NewAppError(err, core.UnexpectedError)
		}
	}

	if !utils.CheckTextHash(retrUser.Password, password) {
		return nil, core.NewAppErrorWithType(core.NotAuthorizedError)
	}

	token, err := utils.CreateJWTToken(*retrUser)
	if err != nil {
		return nil, core.NewAppError(err, core.TokenGenerationError)
	}

	return entities.NewAuthorizedUser(*retrUser, *token), nil
}

func (s *UserService) GetAllUsers() ([]entities.User, error) {
	users, err := s.userRepo.SelectAllUsers()
	if err != nil {
		var coreErr *core.AppError
		wrapped := errors.As(err, &coreErr)
		if wrapped && coreErr.Type == core.NotFoundError {
			return []entities.User{}, nil
		} else {
			return nil, core.NewAppError(err, core.UnexpectedError)
		}
	}

	return users, nil
}

func (s *UserService) FindUserByID(id int) (*entities.User, error) {
	user, err := s.userRepo.SelectUserByID(id)
	if err != nil {
		var coreErr *core.AppError
		wrapped := errors.As(err, &coreErr)
		if (wrapped && coreErr.Type != core.NotFoundError) || !wrapped {
			return nil, core.NewAppError(err, core.UnexpectedError)
		}

		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(userID int, user entities.User) error {
	err := s.userRepo.UpdateUser(userID, user)
	if err != nil {
		var coreErr *core.AppError
		wrapped := errors.As(err, &coreErr)
		if (wrapped && coreErr.Type != core.NotFoundError) || !wrapped {
			return core.NewAppError(err, core.UnexpectedError)
		}

		return err
	}

	return nil
}

func NewUserService(userRepo ports.IUserRepository) *UserService {
	return &UserService{userRepo}
}

package services

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type UserService struct {
	userRepo ports.IUserRepo
}

func (s *UserService) SignUp(user entities.User) (
	*entities.AuthorizedUser,
	error,
) {
	if err := user.ValidateUser(); err != nil {
		return nil, domain.NewAppError(err, domain.ValidationError)
	}

	hash, err := utils.HashText(user.Password)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.HashGenerationError)
	}
	user.Password = hash

	retrUser, err := s.userRepo.InsertUser(user)
	if err != nil {
		return nil, domain.NewAppError(err, domain.UnexpectedError)
	}

	token, err := utils.CreateJWTToken(*retrUser)
	if err != nil {
		return nil, domain.NewAppError(err, domain.TokenGenerationError)
	}

	return entities.NewAuthorizedUser(*retrUser, *token), nil
}

func (s *UserService) SignIn(email, password string) (
	*entities.AuthorizedUser,
	error,
) {
	retrUser, err := s.userRepo.SelectUserByEmail(email)
	if err != nil {
		var coreErr *domain.AppError
		wrapped := errors.As(err, &coreErr)
		if wrapped && coreErr.Type == domain.NotFoundError {
			return nil, domain.NewAppErrorWithType(domain.NotAuthorizedError)
		} else {
			return nil, domain.NewAppError(err, domain.UnexpectedError)
		}
	}

	if !utils.CheckTextHash(retrUser.Password, password) {
		return nil, domain.NewAppErrorWithType(domain.NotAuthorizedError)
	}

	token, err := utils.CreateJWTToken(*retrUser)
	if err != nil {
		return nil, domain.NewAppError(err, domain.TokenGenerationError)
	}

	return entities.NewAuthorizedUser(*retrUser, *token), nil
}

func (s *UserService) GetUsers() ([]entities.User, error) {
	users, err := s.userRepo.SelectUsers()
	if err != nil {
		var coreErr *domain.AppError
		wrapped := errors.As(err, &coreErr)
		if wrapped && coreErr.Type == domain.NotFoundError {
			return []entities.User{}, nil
		} else {
			return nil, domain.NewAppError(err, domain.UnexpectedError)
		}
	}

	return users, nil
}

func (s *UserService) FindUserByID(id int) (*entities.User, error) {
	user, err := s.userRepo.SelectUserByID(id)
	if err != nil {
		var coreErr *domain.AppError
		wrapped := errors.As(err, &coreErr)
		if (wrapped && coreErr.Type != domain.NotFoundError) || !wrapped {
			return nil, domain.NewAppError(err, domain.UnexpectedError)
		}

		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(userID int, user entities.User) error {
	err := s.userRepo.UpdateUser(userID, user)
	if err != nil {
		var coreErr *domain.AppError
		wrapped := errors.As(err, &coreErr)
		if (wrapped && coreErr.Type != domain.NotFoundError) || !wrapped {
			return domain.NewAppError(err, domain.UnexpectedError)
		}

		return err
	}

	return nil
}

func NewUserService(userRepo ports.IUserRepo) *UserService {
	return &UserService{userRepo}
}

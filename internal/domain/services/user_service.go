package services

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
)

// UserService represents a user service implementation.
type UserService struct {
	userRepo ports.IUserRepo
}

func (s *UserService) CreateUser(user entities.User) (*entities.User, error) {
	retrvUser, err := s.userRepo.InsertUser(user)
	if err != nil {
		var appErr *domain.AppError
		converted := errors.As(err, &appErr)
		if !converted {
			return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
		}

		return nil, domain.NewAppError(err, domain.UnexpectedError)
	}

	return retrvUser, nil
}

// GetUsers retrieves a list of users from the system and returns the list
// of users if successful, along with any error encountered during the
// process.
func (s *UserService) GetUsers() ([]entities.User, error) {
	users, err := s.userRepo.SelectUsers()
	if err != nil {
		var appErr *domain.AppError
		converted := errors.As(err, &appErr)
		if converted && appErr.Type == domain.NotFoundError {
			return []entities.User{}, nil
		} else {
			return nil, domain.NewAppError(err, domain.UnexpectedError)
		}
	}

	return users, nil
}

// FindUserByEmail retrieves a user from the system based on the
// provided email and returns the user if found, along with any error
// encountered during the process.
func (s *UserService) FindUserByEmail(email string) (*entities.User, error) {
	user, err := s.userRepo.SelectUserByEmail(email)
	if err != nil {
		var appErr *domain.AppError
		converted := errors.As(err, &appErr)
		if (converted && appErr.Type != domain.NotFoundError) || !converted {
			return nil, domain.NewAppError(err, domain.UnexpectedError)
		}

		return nil, err
	}

	return user, nil
}

// FindUserByID retrieves a user from the system based on the provided user
// ID and returns the user if found, along with any error encountered during
// the process.
func (s *UserService) FindUserByID(id int) (*entities.User, error) {
	user, err := s.userRepo.SelectUserByID(id)
	if err != nil {
		var appErr *domain.AppError
		converted := errors.As(err, &appErr)
		if (converted && appErr.Type != domain.NotFoundError) || !converted {
			return nil, domain.NewAppError(err, domain.UnexpectedError)
		}

		return nil, err
	}

	return user, nil
}

// UpdateUser updates the details of a user in the system based on the
// provided user ID and user object and returns an error, if any,
// encountered during the process.
func (s *UserService) UpdateUser(userID int, user entities.User) error {
	err := s.userRepo.UpdateUser(userID, user)
	if err != nil {
		var appErr *domain.AppError
		converted := errors.As(err, &appErr)
		if (converted && appErr.Type != domain.NotFoundError) || !converted {
			return domain.NewAppError(err, domain.UnexpectedError)
		}

		return err
	}

	return nil
}

func NewUserService(userRepo ports.IUserRepo) *UserService {
	return &UserService{userRepo}
}

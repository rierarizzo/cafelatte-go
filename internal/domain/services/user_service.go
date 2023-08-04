package services

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
)

// UserService represents a user service implementation.
type UserService struct {
	userRepo ports.IUserRepository
}

func (s *UserService) CreateUser(user entities.User) (*entities.User, *domain.AppError) {
	rUser, appErr := s.userRepo.InsertUser(user)
	if appErr != nil {
		return nil, domain.NewAppError(appErr, domain.UnexpectedError)
	}

	return rUser, nil
}

// GetUsers retrieves a list of users from the system and returns the list
// of users if successful, along with any error encountered during the
// process.
func (s *UserService) GetUsers() ([]entities.User, *domain.AppError) {
	rUsers, appErr := s.userRepo.SelectUsers()
	if appErr != nil {
		if appErr.Type == domain.NotFoundError {
			return []entities.User{}, nil
		}

		return nil, appErr
	}

	return rUsers, nil
}

// FindUserByEmail retrieves a user from the system based on the
// provided email and returns the user if found, along with any error
// encountered during the process.
func (s *UserService) FindUserByEmail(email string) (*entities.User, *domain.AppError) {
	user, appErr := s.userRepo.SelectUserByEmail(email)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return user, nil
}

// FindUserByID retrieves a user from the system based on the provided user
// ID and returns the user if found, along with any error encountered during
// the process.
func (s *UserService) FindUserByID(id int) (*entities.User, *domain.AppError) {
	user, appErr := s.userRepo.SelectUserById(id)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return user, nil
}

// UpdateUser updates the details of a user in the system based on the
// provided user ID and user object and returns an error, if any,
// encountered during the process.
func (s *UserService) UpdateUser(userID int,
	user entities.User) *domain.AppError {
	appErr := s.userRepo.UpdateUser(userID, user)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return appErr
	}

	return nil
}

func NewUserService(userRepo ports.IUserRepository) *UserService {
	return &UserService{userRepo}
}

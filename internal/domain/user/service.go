package user

import (
	"fmt"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"mime/multipart"
	"time"
)

type Service struct {
	userRepo      IUserRepository
	userFilesRepo IUserFilesRepository
}

func (s *Service) CreateUser(user User) (*User, *domain.AppError) {
	returnedUser, appErr := s.userRepo.InsertUser(user)
	if appErr != nil {
		return nil, domain.NewAppError(appErr, domain.UnexpectedError)
	}

	return returnedUser, nil
}

// GetUsers retrieves a list of users from the system and returns the list
// of users if successful, along with any error encountered during the
// process.
func (s *Service) GetUsers() ([]User, *domain.AppError) {
	returnedUsers, appErr := s.userRepo.SelectUsers()
	if appErr != nil {
		if appErr.Type == domain.NotFoundError {
			return []User{}, nil
		}

		return nil, appErr
	}

	return returnedUsers, nil
}

// FindUserByEmail retrieves a user from the system based on the
// provided email and returns the user if found, along with any error
// encountered during the process.
func (s *Service) FindUserByEmail(email string) (*User, *domain.AppError) {
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
func (s *Service) FindUserByID(id int) (*User, *domain.AppError) {
	user, appErr := s.userRepo.SelectUserByID(id)
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
func (s *Service) UpdateUser(userID int,
	user User) *domain.AppError {
	appErr := s.userRepo.UpdateUser(userID, user)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return appErr
	}

	return nil
}

func (s *Service) UpdateProfilePic(userID int,
	pic *multipart.FileHeader) (string, *domain.AppError) {
	currentTimeInNano := time.Now().UnixNano()
	picName := fmt.Sprintf("%v-%v", userID, currentTimeInNano)

	picLink, appErr := s.userFilesRepo.UpdateProfilePic(userID, pic, picName)
	if appErr != nil {
		return "", domain.NewAppError(appErr, domain.UnexpectedError)
	}

	return picLink, nil
}

func (s *Service) DeleteUser(userID int) *domain.AppError {
	appErr := s.userRepo.DeleteUser(userID)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return appErr
	}

	return nil
}

func NewUserService(userRepo IUserRepository,
	userFilesRepo IUserFilesRepository) *Service {
	return &Service{
		userRepo:      userRepo,
		userFilesRepo: userFilesRepo,
	}
}

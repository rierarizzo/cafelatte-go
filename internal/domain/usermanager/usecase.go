package usermanager

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/rierarizzo/cafelatte/internal/domain"
)

type DefaultManager struct {
	userRepository   UserRepository
	userFilesStorage UserFilesRepository
}

// GetUsers retrieves a list of users from the system and returns the list
// of users if successful, along with any error encountered during the
// process.
func (manager *DefaultManager) GetUsers() ([]domain.User, *domain.AppError) {
	returnedUsers, appErr := manager.userRepository.SelectUsers()
	if appErr != nil {
		if appErr.Type == domain.NotFoundError {
			return []domain.User{}, nil
		}

		return nil, appErr
	}

	return returnedUsers, nil
}

// FindUserByEmail retrieves a usermanager from the system based on the
// provided email and returns the usermanager if found, along with any error
// encountered during the process.
func (manager *DefaultManager) FindUserByEmail(email string) (
	*domain.User,
	*domain.AppError,
) {
	user, appErr := manager.userRepository.SelectUserByEmail(email)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return user, nil
}

// FindUserById retrieves a usermanager from the system based on the provided usermanager
// ID and returns the usermanager if found, along with any error encountered during
// the process.
func (manager *DefaultManager) FindUserById(id int) (
	*domain.User,
	*domain.AppError,
) {
	user, appErr := manager.userRepository.SelectUserById(id)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return user, nil
}

// UpdateUserById updates the details of a usermanager in the system based on the
// provided usermanager ID and usermanager object and returns an error, if any,
// encountered during the process.
func (manager *DefaultManager) UpdateUserById(
	userID int,
	user domain.User,
) *domain.AppError {
	appErr := manager.userRepository.UpdateUserById(userID, user)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return appErr
	}

	return nil
}

func (manager *DefaultManager) UpdateProfilePicById(
	userID int,
	pic *multipart.FileHeader,
) (string, *domain.AppError) {
	currentTimeInNano := time.Now().UnixNano()
	picName := fmt.Sprintf("%v-%v", userID, currentTimeInNano)

	picLink, appErr := manager.userFilesStorage.UpdateProfilePicById(userID,
		pic, picName)
	if appErr != nil {
		return "", domain.NewAppError(appErr, domain.UnexpectedError)
	}

	return picLink, nil
}

func (manager *DefaultManager) DeleteUserById(userID int) *domain.AppError {
	appErr := manager.userRepository.DeleteUserById(userID)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return appErr
	}

	return nil
}

func New(
	userRepository UserRepository,
	userFilesStorage UserFilesRepository,
) *DefaultManager {
	return &DefaultManager{
		userRepository:   userRepository,
		userFilesStorage: userFilesStorage,
	}
}

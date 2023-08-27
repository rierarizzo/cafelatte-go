package usermanager

import (
	"fmt"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"mime/multipart"
	"time"
)

type DefaultManager struct {
	userRepository   UserRepository
	userFilesStorage UserFilesRepository
}

// GetUsers retrieves a list of users from the system and returns the list
// of users if successful, along with any error encountered during the
// process.
func (m *DefaultManager) GetUsers() ([]domain.User, *domain.AppError) {
	returnedUsers, appErr := m.userRepository.SelectUsers()
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
func (m *DefaultManager) FindUserByEmail(email string) (*domain.User, *domain.AppError) {
	user, appErr := m.userRepository.SelectUserByEmail(email)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return user, nil
}

// FindUserByID retrieves a usermanager from the system based on the provided usermanager
// ID and returns the usermanager if found, along with any error encountered during
// the process.
func (m *DefaultManager) FindUserByID(id int) (*domain.User, *domain.AppError) {
	user, appErr := m.userRepository.SelectUserByID(id)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return user, nil
}

// UpdateUser updates the details of a usermanager in the system based on the
// provided usermanager ID and usermanager object and returns an error, if any,
// encountered during the process.
func (m *DefaultManager) UpdateUser(userID int,
	user domain.User) *domain.AppError {
	appErr := m.userRepository.UpdateUser(userID, user)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return appErr
	}

	return nil
}

func (m *DefaultManager) UpdateProfilePic(userID int,
	pic *multipart.FileHeader) (string, *domain.AppError) {
	currentTimeInNano := time.Now().UnixNano()
	picName := fmt.Sprintf("%v-%v", userID, currentTimeInNano)

	picLink, appErr := m.userFilesStorage.UpdateProfilePic(userID, pic, picName)
	if appErr != nil {
		return "", domain.NewAppError(appErr, domain.UnexpectedError)
	}

	return picLink, nil
}

func (m *DefaultManager) DeleteUser(userID int) *domain.AppError {
	appErr := m.userRepository.DeleteUser(userID)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return appErr
	}

	return nil
}

func New(userRepository UserRepository,
	userFilesStorage UserFilesRepository) *DefaultManager {
	return &DefaultManager{
		userRepository:   userRepository,
		userFilesStorage: userFilesStorage,
	}
}

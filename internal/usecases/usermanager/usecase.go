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

// FindUserById retrieves a usermanager from the system based on the provided usermanager
// Id and returns the usermanager if found, along with any error encountered during
// the process.
func (m *DefaultManager) FindUserById(id int) (*domain.User, *domain.AppError) {
	user, appErr := m.userRepository.SelectUserById(id)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return user, nil
}

// UpdateUserById updates the details of a usermanager in the system based on the
// provided usermanager Id and usermanager object and returns an error, if any,
// encountered during the process.
func (m *DefaultManager) UpdateUserById(userId int, user domain.User) *domain.AppError {
	appErr := m.userRepository.UpdateUserById(userId, user)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return appErr
	}

	return nil
}

func (m *DefaultManager) UpdateProfilePicById(userId int, pic *multipart.FileHeader) (string, *domain.AppError) {
	currentTimeInNano := time.Now().UnixNano()
	picName := fmt.Sprintf("%v-%v", userId, currentTimeInNano)

	picLink, appErr := m.userFilesStorage.UpdateProfilePicById(userId, pic, picName)
	if appErr != nil {
		return "", domain.NewAppError(appErr, domain.UnexpectedError)
	}

	return picLink, nil
}

func (m *DefaultManager) DeleteUserById(userId int) *domain.AppError {
	appErr := m.userRepository.DeleteUserById(userId)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return appErr
	}

	return nil
}

func New(userRepository UserRepository, userFilesStorage UserFilesRepository) *DefaultManager {
	return &DefaultManager{
		userRepository:   userRepository,
		userFilesStorage: userFilesStorage,
	}
}

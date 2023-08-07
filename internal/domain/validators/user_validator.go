package validators

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/constants"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
	"slices"
	"strings"
)

func ValidateUser(user *entities.User) *domain.AppError {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	if appErr := ValidateRole(user); appErr != nil {
		log.Error(appErr)
		return appErr
	}
	if appErr := ValidatePhoneNumber(user); appErr != nil {
		log.Error(appErr)
		return appErr
	}
	if appErr := ValidateEmail(user); appErr != nil {
		log.Error(appErr)
		return appErr
	}

	return nil
}

func ValidateRole(user *entities.User) *domain.AppError {
	if !slices.Contains([]string{"A", "E", "C"}, user.RoleCode) {
		return domain.NewAppError(invalidUserRoleError, domain.ValidationError)
	}

	return nil
}

func ValidatePhoneNumber(user *entities.User) *domain.AppError {
	if len(user.PhoneNumber) != 10 {
		return domain.NewAppError(invalidUserPhoneNumberError,
			domain.ValidationError)
	}

	return nil
}

func ValidateEmail(user *entities.User) *domain.AppError {
	if !strings.Contains(user.Email, "@") {
		return domain.NewAppError(invalidUserEmailError, domain.ValidationError)
	}

	return nil
}

var (
	invalidUserRoleError        = errors.New("invalid role")
	invalidUserPhoneNumberError = errors.New("invalid phone number")
	invalidUserEmailError       = errors.New("invalid email")
)

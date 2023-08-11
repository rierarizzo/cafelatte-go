package user

import (
	"errors"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
	"slices"
	"strings"
)

func ValidateUser(user *User) *domain.AppError {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

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

func ValidateRole(user *User) *domain.AppError {
	if !slices.Contains([]string{"A", "E", "C"}, user.RoleCode) {
		return domain.NewAppError(invalidUserRoleError, domain.ValidationError)
	}

	return nil
}

func ValidatePhoneNumber(user *User) *domain.AppError {
	if len(user.PhoneNumber) != 10 {
		return domain.NewAppError(invalidUserPhoneNumberError,
			domain.ValidationError)
	}

	return nil
}

func ValidateEmail(user *User) *domain.AppError {
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

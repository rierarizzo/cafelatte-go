package authenticator

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
	"slices"
	"strings"
)

func validateUser(user *domain.User) *domain.AppError {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	if appErr := validateUserRole(user); appErr != nil {
		log.Error(appErr)
		return appErr
	}
	if appErr := validateUserPhone(user); appErr != nil {
		log.Error(appErr)
		return appErr
	}
	if appErr := validateUserEmail(user); appErr != nil {
		log.Error(appErr)
		return appErr
	}

	return nil
}

func validateUserRole(user *domain.User) *domain.AppError {
	if !slices.Contains([]string{"A", "E", "C"}, user.RoleCode) {
		return domain.NewAppError(errors.New("invalid usermanager role"),
			domain.ValidationError)
	}

	return nil
}

func validateUserPhone(user *domain.User) *domain.AppError {
	if len(user.PhoneNumber) != 10 {
		return domain.NewAppError(errors.New("invalid usermanager phone"),
			domain.ValidationError)
	}

	return nil
}

func validateUserEmail(user *domain.User) *domain.AppError {
	if !strings.Contains(user.Email, "@") {
		return domain.NewAppError(errors.New("invalid usermanager email"),
			domain.ValidationError)
	}

	return nil
}

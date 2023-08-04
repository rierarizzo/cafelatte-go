package entities

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/constants"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"github.com/sirupsen/logrus"
	"slices"
	"strings"
)

type User struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	PhoneNumber string
	Email       string
	Password    string
	/* A: Admin, E: Employee, C: Client */
	RoleCode     string
	Addresses    []Address
	PaymentCards []PaymentCard
}

func (u *User) HashPassword() *domain.AppError {
	hashed, appErr := utils.HashText(u.Password)
	if appErr != nil {
		return appErr
	}
	u.Password = hashed

	return nil
}

func (u *User) ValidateUser() *domain.AppError {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	if appErr := u.validateRole(); appErr != nil {
		log.Error(appErr)
		return appErr
	}
	if appErr := u.validatePhoneNumber(); appErr != nil {
		log.Error(appErr)
		return appErr
	}
	if appErr := u.validateEmail(); appErr != nil {
		log.Error(appErr)
		return appErr
	}

	return nil
}

func (u *User) validateRole() *domain.AppError {
	if !slices.Contains([]string{"A", "E", "C"}, u.RoleCode) {
		return domain.NewAppError(invalidUserRoleError, domain.ValidationError)
	}

	return nil
}

func (u *User) validatePhoneNumber() *domain.AppError {
	if len(u.PhoneNumber) != 10 {
		return domain.NewAppError(invalidUserPhoneNumberError,
			domain.ValidationError)
	}

	return nil
}

func (u *User) validateEmail() *domain.AppError {
	if !strings.Contains(u.Email, "@") {
		return domain.NewAppError(invalidUserEmailError, domain.ValidationError)
	}

	return nil
}

var (
	invalidUserRoleError        = errors.New("invalid role")
	invalidUserPhoneNumberError = errors.New("invalid phone number")
	invalidUserEmailError       = errors.New("invalid email")
)

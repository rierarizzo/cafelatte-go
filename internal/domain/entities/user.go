package entities

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain/constants"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"github.com/sirupsen/logrus"
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

var (
	invalidUserRoleError        = errors.New("invalid role")
	invalidUserPhoneNumberError = errors.New("invalid phone number")
	invalidUserEmailError       = errors.New("invalid email")
)

func (u *User) HashPassword() error {
	hash, err := utils.HashText(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash

	return nil
}

func (u *User) validateRole() error {
	if u.RoleCode != "A" && u.RoleCode != "E" && u.RoleCode != "C" {
		return invalidUserRoleError
	}

	return nil
}

func (u *User) validatePhoneNumber() error {
	if len(u.PhoneNumber) != 10 {
		return invalidUserPhoneNumberError
	}

	return nil
}

func (u *User) validateEmail() error {
	if !strings.Contains(u.Email, "@") {
		return invalidUserEmailError
	}

	return nil
}

func (u *User) ValidateUser() error {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	if err := u.validateRole(); err != nil {
		log.Error(err)
		return err
	}
	if err := u.validatePhoneNumber(); err != nil {
		log.Error(err)
		return err
	}
	if err := u.validateEmail(); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

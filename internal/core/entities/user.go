package entities

import (
	"github.com/rierarizzo/cafelatte/internal/core/errors"
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
	/* A: Administrador, E: Empleado, C: Cliente */
	RoleCode     string
	Addresses    []Address
	PaymentCards []PaymentCard
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) ValidateUser() error {
	if err := u.validateRole(); err != nil {
		return err
	}

	if err := u.validatePhoneNumber(); err != nil {
		return err
	}

	if err := u.validateEmail(); err != nil {
		return err
	}

	return nil
}

func (u *User) validateRole() error {
	if u.RoleCode != "A" && u.RoleCode != "E" && u.RoleCode != "C" {
		return errors.WrapError(errors.ErrInvalidUserFormat, "role must be 'A', 'E', or 'C'")
	}

	return nil
}

func (u *User) validatePhoneNumber() error {
	if len(u.PhoneNumber) != 10 {
		return errors.WrapError(errors.ErrInvalidUserFormat, "phone number must be 10 digits")
	}

	return nil
}

func (u *User) validateEmail() error {
	if !strings.Contains(u.Email, "@") {
		return errors.WrapError(errors.ErrInvalidUserFormat, "email must contain '@'")
	}

	return nil
}

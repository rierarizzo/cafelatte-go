package entities

import (
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"regexp"
)

type Address struct {
	ID         int
	Type       string
	ProvinceID int
	CityID     int
	PostalCode string
	Detail     string
}

func (a *Address) ValidateAddress() error {
	if err := a.validatePostalCode(); err != nil {
		return err
	}

	return nil
}

func (a *Address) validatePostalCode() error {
	if len(a.PostalCode) != 6 {
		return errors.WrapError(errors.ErrInvalidAddressFormat, "postal code must have 6 digits")
	}

	regex := regexp.MustCompile("^[0-9]+$")

	if !regex.MatchString(a.PostalCode) {
		return errors.WrapError(errors.ErrInvalidAddressFormat, "postal code must have only numbers")
	}

	return nil
}

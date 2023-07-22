package entities

import (
	"fmt"
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
		return fmt.Errorf("%w; postal code must have 6 digits", errors.ErrInvalidAddressFormat)
	}

	regex := regexp.MustCompile("^[0-9]+$")

	if !regex.MatchString(a.PostalCode) {
		return fmt.Errorf("%w; postal code must have only numbers", errors.ErrInvalidAddressFormat)
	}

	return nil
}

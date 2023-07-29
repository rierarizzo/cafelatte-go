package entities

import (
	"errors"
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

var (
	invalidAddressPostalCodeError = errors.New("invalid postal code")
)

func (a *Address) ValidateAddress() error {
	if err := a.validatePostalCode(); err != nil {
		return err
	}

	return nil
}

func (a *Address) validatePostalCode() error {
	if len(a.PostalCode) != 6 {
		return invalidAddressPostalCodeError
	}

	regex := regexp.MustCompile("^[0-9]+$")

	if !regex.MatchString(a.PostalCode) {
		return invalidAddressPostalCodeError
	}

	return nil
}

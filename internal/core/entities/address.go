package entities

import (
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

func (a *Address) IsValidAddress() bool {
	return a.isValidPostalCode()
}

func (a *Address) isValidPostalCode() bool {
	if len(a.PostalCode) != 6 {
		return false
	}

	regex := regexp.MustCompile("^[0-9]+$")

	if !regex.MatchString(a.PostalCode) {
		return false
	}

	return true
}

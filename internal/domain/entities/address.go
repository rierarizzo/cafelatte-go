package entities

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain/constants"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
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

func (a *Address) ValidateAddress() *domain.AppError {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	if appErr := a.validatePostalCode(); appErr != nil {
		log.Error(appErr)
		return appErr
	}

	return nil
}

func (a *Address) validatePostalCode() *domain.AppError {
	if len(a.PostalCode) != 6 {
		return domain.NewAppError(invalidAddressPostalCodeError,
			domain.ValidationError)
	}

	regex := regexp.MustCompile("^[0-9]+$")

	if !regex.MatchString(a.PostalCode) {
		return domain.NewAppError(invalidAddressPostalCodeError,
			domain.ValidationError)
	}

	return nil
}

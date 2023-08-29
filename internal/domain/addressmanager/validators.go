package addressmanager

import (
	"errors"
	"regexp"

	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
)

var invalidAddressPostalCodeError = errors.New("invalid postal code")

func validateAddress(address *domain.Address) *domain.AppError {
	log := logrus.WithField(misc.RequestIdKey, request.Id())

	if appErr := validateAddressPostalCode(address); appErr != nil {
		log.Error(appErr)
		return appErr
	}

	return nil
}

func validateAddressPostalCode(address *domain.Address) *domain.AppError {
	if len(address.PostalCode) != 6 {
		return domain.NewAppError(invalidAddressPostalCodeError,
			domain.ValidationError)
	}

	regex := regexp.MustCompile("^[0-9]+$")

	if !regex.MatchString(address.PostalCode) {
		return domain.NewAppError(invalidAddressPostalCodeError,
			domain.ValidationError)
	}

	return nil
}

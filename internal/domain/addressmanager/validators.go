package addressmanager

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
	"regexp"
)

func ValidateAddress(address *domain.Address) *domain.AppError {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	if appErr := ValidatePostalCode(address); appErr != nil {
		log.Error(appErr)
		return appErr
	}

	return nil
}

func ValidatePostalCode(address *domain.Address) *domain.AppError {
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

var (
	invalidAddressPostalCodeError = errors.New("invalid postal code")
)

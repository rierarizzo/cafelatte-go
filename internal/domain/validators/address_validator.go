package validators

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/pkg/constants"
	"github.com/rierarizzo/cafelatte/pkg/params"
	"github.com/sirupsen/logrus"
	"regexp"
)

func ValidateAddress(address *entities.Address) *domain.AppError {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	if appErr := ValidatePostalCode(address); appErr != nil {
		log.Error(appErr)
		return appErr
	}

	return nil
}

func ValidatePostalCode(address *entities.Address) *domain.AppError {
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

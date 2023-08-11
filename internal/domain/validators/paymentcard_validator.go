package validators

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/pkg/constants"
	"github.com/rierarizzo/cafelatte/pkg/params"
	"github.com/sirupsen/logrus"
	"time"
)

func ValidateType(card *entities.PaymentCard) *domain.AppError {
	if card.Type != "C" && card.Type != "D" {
		return domain.NewAppError(invalidCardTypeError, domain.ValidationError)
	}

	return nil
}

func ValidateCVV(card *entities.PaymentCard) *domain.AppError {
	if len(card.CVV) != 3 && len(card.CVV) != 4 {
		return domain.NewAppError(invalidCardCVVError, domain.ValidationError)
	}

	return nil
}

func ValidateExpirationDateFormat(card *entities.PaymentCard) *domain.AppError {
	if card.ExpirationMonth < 1 || card.ExpirationMonth > 12 {
		return domain.NewAppError(invalidCardExpirationDateError,
			domain.ValidationError)
	}

	return nil
}

func ValidatePaymentCard(card *entities.PaymentCard) *domain.AppError {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	if appErr := ValidateType(card); appErr != nil {
		log.Error(appErr)
		return appErr
	}
	if appErr := ValidateCVV(card); appErr != nil {
		log.Error(appErr)
		return appErr
	}
	if appErr := ValidateExpirationDateFormat(card); appErr != nil {
		log.Error(appErr)
		return appErr
	}

	return nil
}

func ValidateExpirationDate(card *entities.PaymentCard) *domain.AppError {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	expirationDate := time.Date(card.ExpirationYear,
		time.Month(card.ExpirationMonth),
		0, 0, 0, 0, 0, time.UTC)

	if expirationDate.Before(time.Now()) {
		log.Error(expiredCardError)
		return domain.NewAppError(expiredCardError, domain.ValidationError)
	}

	return nil
}

var (
	invalidCardTypeError           = errors.New("invalid card type")
	invalidCardCVVError            = errors.New("invalid CVV")
	invalidCardExpirationDateError = errors.New("invalid expiration date")
	expiredCardError               = errors.New("card is expired")
)

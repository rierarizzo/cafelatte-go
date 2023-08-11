package paymentcard

import (
	"errors"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
	"time"
)

func ValidateType(card *PaymentCard) *domain.AppError {
	if card.Type != "C" && card.Type != "D" {
		return domain.NewAppError(invalidCardTypeError, domain.ValidationError)
	}

	return nil
}

func ValidateCVV(card *PaymentCard) *domain.AppError {
	if len(card.CVV) != 3 && len(card.CVV) != 4 {
		return domain.NewAppError(invalidCardCVVError, domain.ValidationError)
	}

	return nil
}

func ValidateExpirationDateFormat(card *PaymentCard) *domain.AppError {
	if card.ExpirationMonth < 1 || card.ExpirationMonth > 12 {
		return domain.NewAppError(invalidCardExpirationDateError,
			domain.ValidationError)
	}

	return nil
}

func ValidatePaymentCard(card *PaymentCard) *domain.AppError {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

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

func ValidateExpirationDate(card *PaymentCard) *domain.AppError {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

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

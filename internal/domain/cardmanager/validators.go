package cardmanager

import (
	"errors"
	"time"

	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
)

var (
	invalidCardTypeError           = errors.New("invalid card type")
	invalidCardCVVError            = errors.New("invalid CVV")
	invalidCardExpirationDateError = errors.New("invalid expiration date")
	expiredCardError               = errors.New("card is expired")
)

func validateCard(card *domain.PaymentCard) *domain.AppError {
	log := logrus.WithField(misc.RequestIdKey, request.Id())

	if appErr := validateCardType(card); appErr != nil {
		log.Error(appErr)
		return appErr
	}
	if appErr := validateCardCVV(card); appErr != nil {
		log.Error(appErr)
		return appErr
	}
	if appErr := validateCardExpiration(card); appErr != nil {
		return appErr
	}

	return nil
}

func validateCardExpiration(card *domain.PaymentCard) *domain.AppError {
	log := logrus.WithField(misc.RequestIdKey, request.Id())

	if card.ExpirationMonth < 1 || card.ExpirationMonth > 12 {
		return domain.NewAppError(invalidCardExpirationDateError,
			domain.ValidationError)
	}

	expirationDate := time.Date(card.ExpirationYear,
		time.Month(card.ExpirationMonth), 0, 0, 0, 0, 0, time.UTC)

	if expirationDate.Before(time.Now()) {
		log.Error(expiredCardError)
		return domain.NewAppError(expiredCardError, domain.ValidationError)
	}

	return nil
}

func validateCardType(card *domain.PaymentCard) *domain.AppError {
	if card.Type != "C" && card.Type != "D" {
		return domain.NewAppError(invalidCardTypeError, domain.ValidationError)
	}

	return nil
}

func validateCardCVV(card *domain.PaymentCard) *domain.AppError {
	if len(card.CVV) != 3 && len(card.CVV) != 4 {
		return domain.NewAppError(invalidCardCVVError, domain.ValidationError)
	}

	return nil
}

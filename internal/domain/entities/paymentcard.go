package entities

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain/constants"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
	"time"
)

type PaymentCard struct {
	ID int
	/* C: Crédito, D: Débito */
	Type            string
	Company         int
	HolderName      string
	Number          string
	ExpirationYear  int
	ExpirationMonth int
	CVV             string
}

var (
	invalidCardTypeError           = errors.New("invalid card type")
	invalidCardCVVError            = errors.New("invalid CVV")
	invalidCardExpirationDateError = errors.New("invalid expiration date")
	expiredCardError               = errors.New("card is expired")
)

func (c *PaymentCard) validateType() *domain.AppError {
	if c.Type != "C" && c.Type != "D" {
		return domain.NewAppError(invalidCardTypeError, domain.ValidationError)
	}

	return nil
}

func (c *PaymentCard) validateCVV() *domain.AppError {
	if len(c.CVV) != 3 && len(c.CVV) != 4 {
		return domain.NewAppError(invalidCardCVVError, domain.ValidationError)
	}

	return nil
}

func (c *PaymentCard) validateExpirationDateFormat() *domain.AppError {
	if c.ExpirationMonth < 1 || c.ExpirationMonth > 12 {
		return domain.NewAppError(invalidCardExpirationDateError,
			domain.ValidationError)
	}

	return nil
}

func (c *PaymentCard) ValidatePaymentCard() *domain.AppError {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	if appErr := c.validateType(); appErr != nil {
		log.Error(appErr)
		return appErr
	}
	if appErr := c.validateCVV(); appErr != nil {
		log.Error(appErr)
		return appErr
	}
	if appErr := c.validateExpirationDateFormat(); appErr != nil {
		log.Error(appErr)
		return appErr
	}

	return nil
}

func (c *PaymentCard) ValidateExpirationDate() *domain.AppError {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	expirationDate := time.Date(c.ExpirationYear, time.Month(c.ExpirationMonth),
		0, 0, 0, 0, 0, time.UTC)

	if expirationDate.Before(time.Now()) {
		log.Error(expiredCardError)
		return domain.NewAppError(expiredCardError, domain.ValidationError)
	}

	return nil
}

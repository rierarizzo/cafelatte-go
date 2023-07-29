package entities

import (
	"errors"
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

func (c *PaymentCard) validateType() error {
	if c.Type != "C" && c.Type != "D" {
		return invalidCardTypeError
	}

	return nil
}

func (c *PaymentCard) validateCVV() error {
	if len(c.CVV) != 3 && len(c.CVV) != 4 {
		return invalidCardCVVError
	}

	return nil
}

func (c *PaymentCard) validateExpirationDateFormat() error {
	if c.ExpirationMonth < 1 || c.ExpirationMonth > 12 {
		return invalidCardExpirationDateError
	}

	return nil
}

func (c *PaymentCard) ValidatePaymentCard() error {
	if err := c.validateType(); err != nil {
		return err
	}
	if err := c.validateCVV(); err != nil {
		return err
	}
	if err := c.validateExpirationDateFormat(); err != nil {
		return err
	}

	return nil
}

func (c *PaymentCard) ValidateExpirationDate() error {
	expirationDate := time.Date(
		c.ExpirationYear,
		time.Month(c.ExpirationMonth),
		0,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	if expirationDate.Before(time.Now()) {
		return expiredCardError
	}

	return nil
}

package entities

import (
	"fmt"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
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

func (c *PaymentCard) ValidatePaymentCard() error {
	if err := c.validateType(); err != nil {
		return err
	}

	if err := c.validateExpirationDate(); err != nil {
		return err
	}

	if err := c.validateCVV(); err != nil {
		return err
	}

	return nil
}

func (c *PaymentCard) validateType() error {
	if c.Type != "C" && c.Type != "D" {
		return fmt.Errorf("%w; card type must be 'C' or 'D'", errors.ErrInvalidCardFormat)
	}

	return nil
}

func (c *PaymentCard) validateExpirationDate() error {
	if c.ExpirationMonth < 1 || c.ExpirationMonth > 12 {
		return fmt.Errorf("%w; expiration date must be in a valid range [1..12]", errors.ErrInvalidCardFormat)
	}

	year, month, _ := time.Now().Date()

	if c.ExpirationYear < year || (c.ExpirationYear == year && c.ExpirationMonth < int(month)) {
		return errors.ErrExpiredCard
	}

	return nil
}

func (c *PaymentCard) validateCVV() error {
	if len(c.CVV) != 3 && len(c.CVV) != 4 {
		return fmt.Errorf("%w; CVV length must be 3 or 4", errors.ErrInvalidCardFormat)
	}

	return nil
}

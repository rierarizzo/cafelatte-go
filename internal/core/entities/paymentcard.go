package entities

import "time"

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

func (c *PaymentCard) IsValidPaymentCard() bool {
	return c.isValidType() && c.isValidCVV() && c.isValidExpirationDate()
}

func (c *PaymentCard) isValidType() bool {
	return c.Type == "C" || c.Type == "D"
}

func (c *PaymentCard) isValidExpirationDate() bool {
	year, month, _ := time.Now().Date()

	if year < c.ExpirationYear {
		return false
	}

	if year == c.ExpirationYear && int(month) <= c.ExpirationMonth {
		return false
	}

	return true
}

func (c *PaymentCard) isValidCVV() bool {
	return len(c.CVV) == 3 || len(c.CVV) == 4
}

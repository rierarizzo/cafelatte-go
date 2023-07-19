package entities

import "time"

type PaymentCard struct {
	ID int
	/* C: Crédito, D: Débito */
	Type       string
	Number     string
	Expiration time.Time
	CVV        string
}

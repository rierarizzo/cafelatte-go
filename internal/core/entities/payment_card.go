package entities

import "time"

type PaymentCard struct {
	/* C: Crédito, D: Débito */
	Type       string
	Number     string
	Expiration time.Time
	CVV        string
}

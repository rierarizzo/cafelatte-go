package entities

import "time"

type PaymentCard struct {
	ID int
	/* C: Crédito, D: Débito */
	Type           string
	Company        int
	HolderName     string
	Number         string
	ExpirationDate time.Time
	CVV            string
}

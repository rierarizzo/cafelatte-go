package entities

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

package dto

type UserPaymentCardsRequest struct {
	UserID       int                  `json:"userID"`
	PaymentCards []PaymentCardRequest `json:"paymentCards"`
}

type PaymentCardRequest struct {
	Type            string `json:"type"`
	Company         int    `json:"company"`
	HolderName      string `json:"holderName"`
	Number          string `json:"number"`
	ExpirationYear  int    `json:"expirationYear"`
	ExpirationMonth int    `json:"expirationMonth"`
	CVV             string `json:"cvv"`
}

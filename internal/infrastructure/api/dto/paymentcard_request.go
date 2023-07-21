package dto

import (
	"time"
)

type UserPaymentCardsRequest struct {
	UserID       int                  `json:"userID"`
	PaymentCards []PaymentCardRequest `json:"paymentCards"`
}

type PaymentCardRequest struct {
	Type           string    `json:"type"`
	Company        int       `json:"company"`
	HolderName     string    `json:"holderName"`
	Number         string    `json:"number"`
	ExpirationDate time.Time `json:"expirationDate"`
	CVV            string    `json:"cvv"`
}

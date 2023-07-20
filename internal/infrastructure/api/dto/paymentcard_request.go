package dto

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"time"
)

type PaymentCardRequest struct {
	Type           string    `json:"type"`
	Company        int       `json:"company"`
	Issuer         int       `json:"issuer"`
	HolderName     string    `json:"holderName"`
	Number         string    `json:"number"`
	ExpirationDate time.Time `json:"expirationDate"`
	CVV            string    `json:"cvv"`
}

func (pcr *PaymentCardRequest) ToPaymentCardCore() *entities.PaymentCard {
	return &entities.PaymentCard{
		Type:           pcr.Type,
		Company:        pcr.Company,
		Issuer:         pcr.Issuer,
		HolderName:     pcr.HolderName,
		Number:         pcr.Number,
		ExpirationDate: pcr.ExpirationDate,
		CVV:            pcr.CVV,
	}
}

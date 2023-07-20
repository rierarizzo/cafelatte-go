package models

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"time"
)

type PaymentCardModel struct {
	ID             int       `db:"ID"`
	Type           string    `db:"Type"`
	UserID         int       `db:"UserID"`
	Company        int       `db:"Company"`
	Issuer         int       `db:"Issuer"`
	HolderName     string    `db:"HolderName"`
	Number         string    `db:"Number"`
	ExpirationDate time.Time `db:"ExpirationDate"`
	CVV            string    `db:"CVV"`
	Enabled        bool      `db:"Enabled"`
}

func (pcm *PaymentCardModel) LoadFromPaymentCardCore(paymentCard entities.PaymentCard) {
	pcm.ID = paymentCard.ID
	pcm.Type = paymentCard.Type
	pcm.Company = paymentCard.Company
	pcm.Issuer = paymentCard.Issuer
	pcm.HolderName = paymentCard.HolderName
	pcm.Number = paymentCard.Number
	pcm.ExpirationDate = paymentCard.ExpirationDate
	pcm.CVV = paymentCard.CVV
}

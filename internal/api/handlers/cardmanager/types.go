package cardmanager

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CardCreate struct {
	Type            string `json:"type"`
	Company         int    `json:"company"`
	HolderName      string `json:"holderName"`
	Number          string `json:"number"`
	ExpirationYear  int    `json:"expirationYear"`
	ExpirationMonth int    `json:"expirationMonth"`
	CVV             string `json:"cvv"`
}

func (dto CardCreate) Validate() error {
	return validation.ValidateStruct(
		&dto,
		validation.Field(
			&dto.Type, validation.Required, validation.In("C", "D"),
		),
		validation.Field(&dto.Company, validation.Required),
		validation.Field(&dto.HolderName, validation.Required),
		validation.Field(&dto.ExpirationYear, validation.Required),
		validation.Field(&dto.ExpirationMonth, validation.Required),
		validation.Field(&dto.CVV, validation.Required, validation.In(3, 4)),
	)
}

type CardResponse struct {
	Type       string `json:"type"`
	Company    int    `json:"company"`
	HolderName string `json:"holderName"`
}

package dto

type AuthResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"accessToken"`
}

type UserResponse struct {
	ID           int                   `json:"id"`
	CompleteName string                `json:"completeName"`
	Username     string                `json:"username"`
	Email        string                `json:"email"`
	Role         string                `json:"role"`
	Addresses    []AddressResponse     `json:"addresses"`
	PaymentCards []PaymentCardResponse `json:"paymentCards"`
}

type AddressResponse struct {
	Type       string `json:"type"`
	ProvinceID int    `json:"provinceID"`
	CityID     int    `json:"cityID"`
	Detail     string `json:"detail"`
}

type PaymentCardResponse struct {
	Type       string `json:"type"`
	Company    int    `json:"company"`
	HolderName string `json:"holderName"`
}

package dto

type AuthorizedUserResponse struct {
	User        LoggedUserResponse `json:"user"`
	AccessToken string             `json:"accessToken"`
}

type LoggedUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserResponse struct {
	ID           int    `json:"id"`
	CompleteName string `json:"completeName"`
	Username     string `json:"username"`
	PhoneNumber  string `json:"phoneNumber"`
	Email        string `json:"email"`
	Role         string `json:"role"`
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

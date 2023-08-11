package user

type SignUpRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	RoleCode    string `json:"role"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone"`
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

type AddressRequest struct {
	Type       string `json:"type"`
	ProvinceID int    `json:"provinceID"`
	CityID     int    `json:"cityID"`
	PostalCode string `json:"postalCode"`
	Detail     string `json:"detail"`
}

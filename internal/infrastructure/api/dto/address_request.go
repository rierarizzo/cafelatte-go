package dto

type UserAddressesRequest struct {
	UserID    int              `json:"userID"`
	Addresses []AddressRequest `json:"addresses"`
}

type AddressRequest struct {
	Type       string `json:"type"`
	ProvinceID int    `json:"provinceID"`
	CityID     int    `json:"cityID"`
	PostalCode string `json:"postalCode"`
	Detail     string `json:"detail"`
}

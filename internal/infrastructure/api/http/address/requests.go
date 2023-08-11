package address

type CreateRequest struct {
	Type       string `json:"type"`
	ProvinceID int    `json:"provinceID"`
	CityID     int    `json:"cityID"`
	PostalCode string `json:"postalCode"`
	Detail     string `json:"detail"`
}

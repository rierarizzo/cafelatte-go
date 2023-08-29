package addressmanager

type RegisterAddressRequest struct {
	Type       string `json:"type"`
	ProvinceId int    `json:"provinceId"`
	CityId     int    `json:"cityId"`
	PostalCode string `json:"postalCode"`
	Detail     string `json:"detail"`
}

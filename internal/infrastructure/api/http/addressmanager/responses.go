package addressmanager

type AddressResponse struct {
	Type       string `json:"type"`
	ProvinceId int    `json:"provinceId"`
	CityId     int    `json:"cityId"`
	Detail     string `json:"detail"`
}

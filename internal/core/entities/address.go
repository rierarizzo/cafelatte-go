package entities

type Address struct {
	ID           int
	ProvinceID   int
	ProvinceName string
	CityID       int
	CityName     string
	PostalCode   string
	Detail       string
}

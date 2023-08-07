package entities

type Address struct {
	ID         int
	Type       string
	ProvinceID int
	CityID     int
	PostalCode string
	Detail     string
}

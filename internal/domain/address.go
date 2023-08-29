package domain

type Address struct {
	Id         int
	Type       string
	ProvinceId int
	CityId     int
	PostalCode string
	Detail     string
}

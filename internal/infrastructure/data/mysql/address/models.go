package address

import (
	"database/sql"
	"time"
)

type Model struct {
	Id         sql.NullInt64 `db:"Id"`
	Type       string        `db:"Type"`
	UserId     int           `db:"UserId"`
	ProvinceId int           `db:"ProvinceId"`
	CityId     int           `db:"CityId"`
	PostalCode string        `db:"PostalCode"`
	Detail     string        `db:"Detail"`
	Status     bool          `db:"Status"`
	CreatedAt  time.Time     `db:"CreatedAt"`
	UpdatedAt  time.Time     `db:"UpdatedAt"`
}

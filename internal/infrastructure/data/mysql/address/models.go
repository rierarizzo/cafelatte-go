package address

import (
	"database/sql"
	"time"
)

type Model struct {
	ID         sql.NullInt64 `db:"ID"`
	Type       string        `db:"Type"`
	UserID     int           `db:"UserID"`
	ProvinceID int           `db:"ProvinceID"`
	CityID     int           `db:"CityID"`
	PostalCode string        `db:"PostalCode"`
	Detail     string        `db:"Detail"`
	Status     bool          `db:"Status"`
	CreatedAt  time.Time     `db:"CreatedAt"`
	UpdatedAt  time.Time     `db:"UpdatedAt"`
}
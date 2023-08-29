package card

import (
	"database/sql"
	"time"
)

type Model struct {
	Id              sql.NullInt64 `db:"Id"`
	Type            string        `db:"Type"`
	UserId          int           `db:"UserId"`
	Company         int           `db:"Company"`
	HolderName      string        `db:"HolderName"`
	Number          string        `db:"Number"`
	ExpirationYear  int           `db:"ExpirationYear"`
	ExpirationMonth int           `db:"ExpirationMonth"`
	CVV             string        `db:"CVV"`
	Status          bool          `db:"Status"`
	CreatedAt       time.Time     `db:"CreatedAt"`
	UpdatedAt       time.Time     `db:"UpdatedAt"`
}

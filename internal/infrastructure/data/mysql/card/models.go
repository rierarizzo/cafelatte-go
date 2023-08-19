package card

import (
	"database/sql"
	"time"
)

type Model struct {
	ID              sql.NullInt64 `db:"ID"`
	Type            string        `db:"Type"`
	UserID          int           `db:"UserID"`
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

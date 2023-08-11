package product

import (
	"time"
)

type Model struct {
	ID           int       `db:"ID"`
	Name         string    `db:"Name"`
	Description  string    `db:"Description"`
	ImageURL     string    `db:"ImageURL"`
	Price        float64   `db:"Price"`
	CategoryCode string    `db:"CategoryCode"`
	Stock        int       `db:"Stock"`
	Status       bool      `db:"Status"`
	CreatedAt    time.Time `db:"CreatedAt"`
	UpdatedAt    time.Time `db:"UpdatedAt"`
}

type CategoryModel struct {
	Code        string `db:"Code"`
	Description string `db:"Description"`
}

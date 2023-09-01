package product

import (
	"time"
)

type Model struct {
	Id           int       `db:"Id"`
	Name         string    `db:"Name"`
	Description  string    `db:"Description"`
	ImageUrl     string    `db:"ImageUrl"`
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

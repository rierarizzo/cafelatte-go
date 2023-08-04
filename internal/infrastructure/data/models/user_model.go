package models

import (
	"time"
)

type UserModel struct {
	ID          int       `db:"ID"`
	Username    string    `db:"Username"`
	Name        string    `db:"Name"`
	Surname     string    `db:"Surname"`
	PhoneNumber string    `db:"PhoneNumber"`
	Email       string    `db:"Email"`
	Password    string    `db:"Password"`
	RoleCode    string    `db:"RoleCode"`
	Status      bool      `db:"Status"`
	CreatedAt   time.Time `db:"CreatedAt"`
	UpdatedAt   time.Time `db:"UpdatedAt"`
}

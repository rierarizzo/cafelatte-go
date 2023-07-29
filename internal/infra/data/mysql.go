package data

import (
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect(dsn string) *sqlx.DB {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	return db
}

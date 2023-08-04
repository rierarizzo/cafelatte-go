package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func Connect(dsn string) *sqlx.DB {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		logrus.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		logrus.Panic(err)
	}

	return db
}

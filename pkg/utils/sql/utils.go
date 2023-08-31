package sql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/sirupsen/logrus"
)

func StartTransaction(db *sqlx.DB) (*sqlx.Tx, *domain.AppError) {
	tx, err := db.Beginx()
	if err != nil {
		logrus.Error(err)
		return nil, domain.NewAppError(err, domain.RepositoryError)
	}
	return tx, nil
}

func CommitTransaction(tx *sqlx.Tx) *domain.AppError {
	if err := tx.Commit(); err != nil {
		return domain.NewAppError(err, domain.RepositoryError)
	}

	return nil
}

func RollbackIfPanic(tx *sqlx.Tx) {
	if p := recover(); p != nil {
		_ = tx.Rollback()
	}
}

func ExecWithTransaction(tx *sqlx.Tx, query string,
	args ...interface{}) (sql.Result, *domain.AppError) {
	result, err := tx.Exec(query, args...)
	if err != nil {
		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	return result, nil
}

func GetLastInsertedId(result sql.Result) (int, *domain.AppError) {
	lastId, err := result.LastInsertId()
	if err != nil {
		appErr := domain.NewAppError(err, domain.RepositoryError)
		return 0, appErr
	}

	return int(lastId), nil
}

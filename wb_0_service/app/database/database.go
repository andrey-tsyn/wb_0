package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDb(driverName, connStr string) (*sqlx.DB, error) {
	sqlxDb, err := sqlx.Connect(driverName, connStr)
	if err != nil {
		return nil, err
	}
	err = sqlxDb.Ping()
	if err != nil {
		return nil, err
	}

	return sqlxDb, nil
}

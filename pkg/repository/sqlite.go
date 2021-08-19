package repository

import (
	"database/sql"
)

type Config struct {
	DBName string
}

func NewSqliteDB(cfg Config) (*sql.DB, error)  {
	db, err := sql.Open("sqlite3", cfg.DBName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}


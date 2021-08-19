package repository

import (
	"database/sql"
)

type Account interface {
	GetCurrency() (string, error)
	GetCash() (float64, error)
	UpdateCash(float64) error
}

type Repository struct {
	Account
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Account: NewAccountSqlite(db),
	}
}
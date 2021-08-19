package repository

import (
	"database/sql"
)

type AccountSqlite struct {
	db *sql.DB
}

const MyID = 21

func NewAccountSqlite(db *sql.DB) *AccountSqlite {
	return &AccountSqlite{db: db}
}

func (r *AccountSqlite) GetCurrency() (string,error) {
	var cur string
	err := r.db.QueryRow("SELECT currency FROM bank WHERE id = $1", MyID).Scan(&cur)
	if err != nil {
		return "", err
	}
	return cur, nil
}

func (r *AccountSqlite) GetCash() (float64,error) {
	var cash float64
	err := r.db.QueryRow("SELECT cash FROM bank WHERE id = $1", MyID).Scan(&cash)
	if err != nil {
		return 0, err
	}
	return cash, nil
}

func (r *AccountSqlite) UpdateCash(sum float64) error {
	_, err := r.db.Exec("UPDATE bank SET cash = $1 WHERE id = $2", sum, MyID)
	if err != nil {
		return err
	}

	return nil
}



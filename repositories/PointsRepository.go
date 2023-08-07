package repositories

import (
	"database/sql"
)

func GetClienteSaldo(db *sql.DB, documento string) (float64, error) {
	var saldo float64
	err := db.QueryRow("SELECT SUM(valor) as saldo FROM points WHERE documento = ?", documento).Scan(&saldo)
	return saldo, err
}

func GetExtrato(db *sql.DB, documento string) (*sql.Rows, error) {
	rows, err := db.Query("SELECT * FROM points WHERE documento = ?", documento)
	return rows, err
}

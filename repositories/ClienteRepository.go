package repositories

import (
	"database/sql"
)

func GetClientTotal(db *sql.DB, documento string) (int, error) {
	var total int
	err := db.QueryRow("SELECT COUNT(*) as total FROM clientes WHERE documento = ?", documento).Scan(&total)
	return total, err
}

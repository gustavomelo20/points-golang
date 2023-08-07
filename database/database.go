package database

import (
	"database/sql"
	"fmt"
	"log"
)

func Conect(db *sql.DB) error {

	dsn := "root:points@tcp(localhost:8893)/go_db"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida.")

	return nil
}

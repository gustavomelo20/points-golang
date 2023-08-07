package main

import (
	"crud/clientes"
	"crud/helpers"
	"crud/points"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var db *sql.DB

func main() {

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

	http.HandleFunc("/cliente", handleCliente)
	http.HandleFunc("/atribuir", handleAtribuir)
	http.HandleFunc("/trocar", handleTrocar)
	http.HandleFunc("/extrato", handleExtrato)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCliente(w http.ResponseWriter, r *http.Request) {

	helpers.HeadersCors(w, r)

	switch r.Method {
	case http.MethodGet:
		clientes.ListaCliente(w, r, db)
	case http.MethodPost:
		clientes.CriarCliente(w, r, db)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func handleAtribuir(w http.ResponseWriter, r *http.Request) {

	helpers.HeadersCors(w, r)

	switch r.Method {
	case http.MethodPost:
		points.CreditarPoints(w, r, db)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func handleTrocar(w http.ResponseWriter, r *http.Request) {

	helpers.HeadersCors(w, r)

	switch r.Method {
	case http.MethodPost:
		points.DebitarPoints(w, r, db)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func handleExtrato(w http.ResponseWriter, r *http.Request) {

	helpers.HeadersCors(w, r)

	switch r.Method {
	case http.MethodGet:
		points.ExtratoPoints(w, r, db)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

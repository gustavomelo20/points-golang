package main

import (
	"crud/clientes"
	"crud/points"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

// Cliente representa a estrutura do cliente

var db *sql.DB

func main() {
	// Substitua 'seu_usuario' e 'sua_senha' pelas credenciais corretas do banco de dados
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

	http.HandleFunc("/cliente", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			clientes.ListaCliente(w, r, db)
		case http.MethodPost:
			clientes.CriarCliente(w, r, db)
		default:
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/atribuir", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			points.CreditarPoints(w, r, db)
		default:
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/trocar", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			points.DebitarPoints(w, r, db)
		default:
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/extrato", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			points.ExtratoPoints(w, r, db)
		default:
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

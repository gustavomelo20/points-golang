package clientes

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func ListaCliente(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT * FROM clientes")
	if err != nil {
		http.Error(w, "Falha ao listar clientes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	clientes := []Cliente{}

	for rows.Next() {
		var cliente Cliente
		err := rows.Scan(&cliente.ID, &cliente.Nome, &cliente.Email, &cliente.Telefone, &cliente.Documento)
		if err != nil {
			http.Error(w, "Falha ao ler dados do cliente", http.StatusInternalServerError)
			return
		}
		clientes = append(clientes, cliente)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Falha ao listar clientes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clientes)
}

func CriarCliente(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var novoCliente Cliente

	err := json.NewDecoder(r.Body).Decode(&novoCliente)
	if err != nil {
		http.Error(w, "Falha ao ler dados do cliente", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO clientes (tipo, valor, telefone, documento) VALUES (?, ?, ?, ?)", novoCliente.Nome, novoCliente.Email, novoCliente.Telefone, novoCliente.Documento)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok && driverErr.Number == 1062 {
			http.Error(w, "Cliente com este e-mail já existe.", http.StatusConflict)
			return
		}
		http.Error(w, "Falha ao criar cliente", http.StatusInternalServerError)
		return
	}

	lastInsertID, _ := result.LastInsertId()
	novoCliente.ID = int(lastInsertID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoCliente)
}

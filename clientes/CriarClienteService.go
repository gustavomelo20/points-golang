package clientes

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func CriarCliente(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var novoCliente Cliente

	err := json.NewDecoder(r.Body).Decode(&novoCliente)
	if err != nil {
		http.Error(w, "Falha ao ler dados do cliente", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO clientes (nome, email, telefone, documento) VALUES (?, ?, ?, ?)", novoCliente.Nome, novoCliente.Email, novoCliente.Telefone, novoCliente.Documento)
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

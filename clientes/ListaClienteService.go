package clientes

import (
	"database/sql"
	"encoding/json"
	"net/http"
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

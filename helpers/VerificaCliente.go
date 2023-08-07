package helpers

import (
	"crud/repositories"
	"database/sql"
	"fmt"
	"net/http"
)

func VerificaCliente(w http.ResponseWriter, db *sql.DB, documento string) error {
	total, err := repositories.GetClientTotal(db, documento)
	if err != nil {
		http.Error(w, "Falha ao buscar cliente", http.StatusInternalServerError)
		return err
	}

	if total == 0 {
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return fmt.Errorf("Cliente não encontrado")
	}

	return nil
}

func HeadersCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
}

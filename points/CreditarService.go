package points

import (
	"crud/helpers"
	"database/sql"
	"encoding/json"
	"net/http"
)

func CreditarPoints(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var novoPoint Points
	err := json.NewDecoder(r.Body).Decode(&novoPoint)
	if err != nil {
		http.Error(w, "Falha ao ler json", http.StatusBadRequest)
		return
	}

	if err := helpers.VerificaCliente(w, db, novoPoint.Documento); err != nil {
		return
	}

	tipo := "credito"

	result, err := db.Exec("INSERT INTO points (tipo, valor, documento) VALUES (?, ?, ?)", tipo, novoPoint.Valor, novoPoint.Documento)
	if err != nil {
		http.Error(w, "Falha ao criar point", http.StatusInternalServerError)
		return
	}

	lastInsertID, _ := result.LastInsertId()
	novoPoint.ID = int(lastInsertID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoPoint)
}

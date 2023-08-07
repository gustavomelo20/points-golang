package points

import (
	"crud/helpers"
	"crud/repositories"
	"database/sql"
	"encoding/json"
	"net/http"
)

func DebitarPoints(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var novoPoint Points
	err := json.NewDecoder(r.Body).Decode(&novoPoint)
	if err != nil {
		http.Error(w, "Falha ao ler json", http.StatusBadRequest)
		return
	}

	if err := helpers.VerificaCliente(w, db, novoPoint.Documento); err != nil {
		return
	}

	saldo, err := repositories.GetClienteSaldo(db, novoPoint.Documento)
	if err != nil {
		http.Error(w, "Falha ao buscar saldo", http.StatusInternalServerError)
		return
	}

	if novoPoint.Valor > saldo {
		http.Error(w, "Ta chapando? Você não tem essa bala, fio!", http.StatusNotFound)
		return
	}

	tipo := "debito"

	result, err := db.Exec("INSERT INTO points (tipo, valor, documento) VALUES (?, ?, ?)", tipo, novoPoint.Valor*-1, novoPoint.Documento)
	if err != nil {
		http.Error(w, "Falha ao criar point", http.StatusInternalServerError)
		return
	}

	novoPoint.Tipo = tipo

	lastInsertID, _ := result.LastInsertId()
	novoPoint.ID = int(lastInsertID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoPoint)
}

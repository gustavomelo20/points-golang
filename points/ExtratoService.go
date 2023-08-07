package points

import (
	"crud/repositories"
	"database/sql"
	"encoding/json"
	"net/http"
)

func ExtratoPoints(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var novoPoint Points
	documento := r.URL.Query().Get("documento")
	if documento == "" {
		http.Error(w, "Documento não fornecido na URL", http.StatusBadRequest)
		return
	}

	novoPoint.Documento = documento

	rows, err := repositories.GetExtrato(db, novoPoint.Documento)
	if err != nil {
		http.Error(w, "Erro ao buscar dados no banco de dados", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var extrato []Points
	for rows.Next() {
		var p Points
		err := rows.Scan(&p.ID, &p.Tipo, &p.Valor, &p.Documento)
		if err != nil {
			http.Error(w, "Erro ao ler dados do banco de dados", http.StatusInternalServerError)
			return
		}
		extrato = append(extrato, p)
	}

	saldo, err := repositories.GetClienteSaldo(db, novoPoint.Documento)

	response := Response{
		Extrato: extrato,
		Saldo:   saldo,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

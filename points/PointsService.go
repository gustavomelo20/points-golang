package points

import (
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

	var total int
	err = db.QueryRow("SELECT COUNT(*) as total FROM clientes WHERE documento = ?", novoPoint.Documento).Scan(&total)
	if total == 0 {
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
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

func DebitarPoints(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var novoPoint Points
	err := json.NewDecoder(r.Body).Decode(&novoPoint)
	if err != nil {
		http.Error(w, "Falha ao ler json", http.StatusBadRequest)
		return
	}

	var total int
	err = db.QueryRow("SELECT COUNT(*) as total FROM clientes WHERE documento = ?", novoPoint.Documento).Scan(&total)
	if total == 0 {
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return
	}

	var saldo float64
	err = db.QueryRow("SELECT SUM(valor) as saldo FROM points WHERE documento = ?", novoPoint.Documento).Scan(&saldo)
	if novoPoint.Valor > saldo {
		http.Error(w, "Ta chapando ? ce não tem essa bala fio!", http.StatusNotFound)
		return
	}

	tipo := "debito"

	result, err := db.Exec("INSERT INTO points (tipo, valor, documento) VALUES (?, ?, ?)", tipo, novoPoint.Valor*-1, novoPoint.Documento)
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

func ExtratoPoints(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var novoPoint Points
	err := json.NewDecoder(r.Body).Decode(&novoPoint)
	if err != nil {
		http.Error(w, "Falha ao ler json", http.StatusBadRequest)
		return
	}

	rows, err := db.Query("SELECT * FROM points WHERE documento = ?", novoPoint.Documento)
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

	var saldo float64
	err = db.QueryRow("SELECT SUM(Valor) as saldo FROM points WHERE documento = ?", novoPoint.Documento).Scan(&saldo)

	type Response struct {
		Saldo   float64  `json:"saldo"`
		Extrato []Points `json:"Extrato"`
	}

	response := Response{
		Extrato: extrato,
		Saldo:   saldo,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

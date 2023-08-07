package points

type Points struct {
	ID        int     `json:"id"`
	Valor     float64 `json:"valor"`
	Saldo     float64 `json:"saldo"`
	Tipo      string  `json:"tipo"`
	Documento string  `json:"documento"`
}

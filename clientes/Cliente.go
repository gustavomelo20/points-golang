package clientes

type Cliente struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	Documento string `json:"documento"`
	Telefone  string `json:"telefone"`
}

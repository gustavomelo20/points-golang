package points

type Response struct {
	Saldo   float64  `json:"saldo"`
	Extrato []Points `json:"Extrato"`
}

package model

// Saldo - struct
type Saldo struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Saldo float64 `json:"saldo"`
	} `json:"dados"`
}

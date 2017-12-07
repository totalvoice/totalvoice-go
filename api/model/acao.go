package model

// Acao struct
type Acao struct {
	Acao      string `json:"acao"`
	Opcao     string `json:"opcao"`
	Menu      string `json:"menu"`
	AcaoDados interface {
	} `json:"acao_dados"`
}

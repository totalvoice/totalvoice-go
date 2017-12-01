package model

// Acao struct
type Acao struct {
	Acao      string `json:"acao"`
	AcaoDados interface {
	} `json:"acao_dados"`
}

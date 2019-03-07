package model

import (
	"time"
)

// ValidaNumero model
type ValidaNumero struct {
	NumeroDestino	string	`json:"numero_destino"`
}

// ValidaNumeroResponse model
type ValidaNumeroResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		ID            int         `json:"id"`
		NumeroDestino string	  `json:"numero_destino"`
		DataCriacao   time.Time   `json:"data_criacao"`
		Preco         float64	  `json:"preco"`
		Valido        bool        `json:"valido"`
		Finalizado    bool        `json:"finalizado"`
	} `json:"dados"`
}
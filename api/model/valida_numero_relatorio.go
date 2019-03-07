package model

import (
	"time"
)

// ValidaNumeroRelatorio struct
type ValidaNumeroRelatorio struct{}

// ValidaNumeroRelatorioResponse struct
type ValidaNumeroRelatorioResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Relatorio []struct {
			ID            int         `json:"id"`
			NumeroDestino string	  `json:"numero_destino"`
			DataCriacao   time.Time   `json:"data_criacao"`
			Preco         float64	  `json:"preco"`
			Valido        bool        `json:"valido"`
			Finalizado    bool        `json:"finalizado"`
		} `json:"relatorio"`
	} `json:"dados"`
}

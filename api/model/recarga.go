package model

import "time"

// Recarga -
type Recarga struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Relatorio []struct {
			ID        int       `json:"id"`
			Credito   float64   `json:"credito"`
			Data      time.Time `json:"data"`
			Descricao string    `json:"descricao"`
		} `json:"relatorio"`
	} `json:"dados"`
}

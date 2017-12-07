package model

import (
	"time"
)

// SMSRelatorio struct
type SMSRelatorio struct {
	DataInicio string `json:"data_inicio"`
	DataFim    string `json:"data_fim"`
}

// SMSRelatorioResponse struct
type SMSRelatorioResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Relatorio []struct {
			ID              int           `json:"id"`
			NumeroDestino   string        `json:"numero_destino"`
			DataCriacao     time.Time     `json:"data_criacao"`
			DataEnvio       interface{}   `json:"data_envio"`
			Mensagem        string        `json:"mensagem"`
			Preco           int           `json:"preco"`
			Status          string        `json:"status"`
			RespostaUsuario bool          `json:"resposta_usuario"`
			Respostas       []interface{} `json:"respostas"`
		} `json:"relatorio"`
	} `json:"dados"`
}

package model

import "time"

// SMS model
type SMS struct {
	NumeroDestino   string      `json:"numero_destino"`
	Mensagem        string      `json:"mensagem"`
	RespostaUsuario bool        `json:"resposta_usuario"`
	MultiSMS        bool        `json:"multi_sms"`
	DataCriacao     interface{} `json:"data_criacao"`
}

// SMSResponse struct
type SMSResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		ID              int         `json:"id"`
		NumeroDestino   string      `json:"numero_destino"`
		DataCriacao     time.Time   `json:"data_criacao"`
		DataEnvio       interface{} `json:"data_envio"`
		Mensagem        string      `json:"mensagem"`
		Preco           float64     `json:"preco"`
		Status          string      `json:"status"`
		RespostaUsuario bool        `json:"resposta_usuario"`
		Respostas       []struct {
			ID           int    `json:"id"`
			Resposta     string `json:"resposta"`
			DataResposta string `json:"data_resposta"`
		} `json:"respostas"`
	} `json:"dados"`
}

package model

import (
	"time"
)

// Bina model
type Bina struct {
	Telefone string `json:"telefone"`
}

// ValidoResponse model
type ValidoResponse struct {
	Valido bool `json:"valido"`
}

// BinaResponse struct
type BinaResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		ID                     int       `json:"id"`
		NumeroDestino          string    `json:"numero_destino"`
		DataCriacao            time.Time `json:"data_criacao"`
		DataInicio             time.Time `json:"data_inicio"`
		Tipo                   string    `json:"tipo"`
		Status                 string    `json:"status"`
		DuracaoSegundos        int       `json:"duracao_segundos"`
		Duracao                string    `json:"duracao"`
		DuracaoCobradaSegundos int       `json:"duracao_cobrada_segundos"`
		DuracaoCobrada         string    `json:"duracao_cobrada"`
		DuracaoFaladaSegundos  int       `json:"duracao_falada_segundos"`
		DuracaoFalada          string    `json:"duracao_falada"`
		Preco                  int       `json:"preco"`
		URLAudio               string    `json:"url_audio"`
		RespostaUsuario        bool      `json:"resposta_usuario"`
		Resposta               string    `json:"resposta"`
	} `json:"dados"`
}

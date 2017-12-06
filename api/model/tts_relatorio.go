package model

import (
	"time"
)

// TTSRelatorio struct
type TTSRelatorio struct {
	DataInicio string `json:"data_inicio"`
	DataFim    string `json:"data_fim"`
}

// TTSRelatorioResponse struct
type TTSRelatorioResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		Relatorio []struct {
			ID                     int         `json:"id"`
			NumeroDestino          string      `json:"numero_destino"`
			DataCriacao            time.Time   `json:"data_criacao"`
			DataInicio             time.Time   `json:"data_inicio"`
			Tipo                   string      `json:"tipo"`
			Status                 string      `json:"status"`
			DuracaoSegundos        int         `json:"duracao_segundos"`
			Duracao                string      `json:"duracao"`
			DuracaoCobradaSegundos int         `json:"duracao_cobrada_segundos"`
			DuracaoCobrada         string      `json:"duracao_cobrada"`
			DuracaoFaladaSegundos  interface{} `json:"duracao_falada_segundos"`
			DuracaoFalada          interface{} `json:"duracao_falada"`
			Preco                  float64     `json:"preco"`
			Mensagem               string      `json:"mensagem"`
			Velocidade             int         `json:"velocidade"`
			RespostaUsuario        bool        `json:"resposta_usuario"`
			Resposta               string      `json:"resposta"`
		} `json:"relatorio"`
	} `json:"dados"`
}

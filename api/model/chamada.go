package model

import "time"

// Chamada struct
type Chamada struct {
	NumeroOrigem  string    `json:"numero_origem"`
	NumeroDestino string    `json:"numero_destino"`
	DataCriacao   time.Time `json:"data_criacao"`
	GravarAudio   bool      `json:"gravar_audio"`
	BinaOrigem    string    `json:"bina_origem"`
	BinaDestino   string    `json:"bina_destino"`
	Tags          string    `json:"tags"`
}

// ChamadaResponse struct
type ChamadaResponse struct {
	Status   int    `json:"status"`
	Sucesso  bool   `json:"sucesso"`
	Motivo   int    `json:"motivo"`
	Mensagem string `json:"mensagem"`
	Dados    struct {
		ID            int         `json:"id"`
		DataCriacao   time.Time   `json:"data_criacao"`
		Ativa         bool        `json:"ativa"`
		ClienteID     int         `json:"cliente_id"`
		RamalIDOrigem interface{} `json:"ramal_id_origem"`
		Origem        struct {
			DataInicio             time.Time `json:"data_inicio"`
			Numero                 string    `json:"numero"`
			Tipo                   string    `json:"tipo"`
			Status                 string    `json:"status"`
			Substatus              string    `json:"substatus"`
			DuracaoSegundos        int       `json:"duracao_segundos"`
			Duracao                string    `json:"duracao"`
			DuracaoCobradaSegundos int       `json:"duracao_cobrada_segundos"`
			DuracaoCobrada         string    `json:"duracao_cobrada"`
			DuracaoFaladaSegundos  int       `json:"duracao_falada_segundos"`
			DuracaoFalada          string    `json:"duracao_falada"`
			Preco                  float64   `json:"preco"`
		} `json:"origem"`
		Destino struct {
			DataInicio             time.Time `json:"data_inicio"`
			Numero                 string    `json:"numero"`
			Tipo                   string    `json:"tipo"`
			Status                 string    `json:"status"`
			DuracaoSegundos        int       `json:"duracao_segundos"`
			Duracao                string    `json:"duracao"`
			DuracaoCobradaSegundos int       `json:"duracao_cobrada_segundos"`
			DuracaoCobrada         string    `json:"duracao_cobrada"`
			DuracaoFaladaSegundos  int       `json:"duracao_falada_segundos"`
			DuracaoFalada          string    `json:"duracao_falada"`
			Preco                  float64   `json:"preco"`
		} `json:"destino"`
	} `json:"dados"`
}

package model

import (
	"time"
)

// Composto struct
type Composto struct {
	NumeroDestino string `json:"numero_destino"`
	Dados  []Acao `json:"dados"`
	Bina   string `json:"bina"`
	Tags   string `json:"tags"`
}

// AddAcao - Adiciona uma ação no JSON
func (comp *Composto) AddAcao(acao Acao) []Acao {
	comp.Dados = append(comp.Dados, acao)
	return comp.Dados
}

// CompostoResponse struct
type CompostoResponse struct {
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

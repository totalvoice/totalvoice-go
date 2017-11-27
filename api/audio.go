package api

import (
	"strconv"
	"time"
)

// Audio struct
type Audio struct {
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

// AudioService service
type AudioService struct {
	Client HTTPClient

	Relatorio *AudioRelatorioService
}

// Enviar - Envia uma mensagem de audio
func (s AudioService) Enviar(numero string, urlAudio string, respostaUsuario bool, bina string) (*Audio, error) {

	params := map[string]string{
		"numero_destino":   numero,
		"url_audio":        urlAudio,
		"resposta_usuario": strconv.FormatBool(respostaUsuario),
		"bina":             bina,
	}

	resp := new(Audio)
	err := s.Client.CreateResource(params, RotaAudio, resp)
	return resp, err
}

// BuscaAudio - Busca uma mensagem de audio pelo seu ID
func (s AudioService) BuscaAudio(id int) (*Audio, error) {
	sID := strconv.Itoa(id)
	resp := new(Audio)
	err := s.Client.GetResource(RotaAudio, sID, resp)
	return resp, err
}
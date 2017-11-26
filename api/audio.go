package api

import (
	"strconv"
	"time"

	"github.com/totalvoice/go-client/handler"
)

const (
	// RotaAudio - rota para manipulação de audios na API
	RotaAudio = "/audio"
)

// AudioService service
type AudioService struct {
	Client HTTPClient
}

// Enviar - Envia uma mensagem de audio
func (s AudioService) Enviar(numero string, urlAudio string, respostaUsuario bool, bina string) (string, error) {

	params := map[string]string{
		"numero_destino":   numero,
		"url_audio":        urlAudio,
		"resposta_usuario": strconv.FormatBool(respostaUsuario),
		"bina":             bina,
	}

	err := s.Client.CreateResource(params, RotaAudio)
	return "", err
}

// BuscaAudio - Busca uma mensagem de audio pelo seu ID
func (s AudioService) BuscaAudio(id int) (string, error) {
	sID := strconv.Itoa(id)
	err := s.Client.GetResource(RotaAudio, sID, nil)
	return "", err
}

// Relatorio - Relatório de mensagens de Audio
func (s AudioService) Relatorio(dataInicial time.Time, dataFinal time.Time) (handler.TvceResponse, error) {

	params := map[string]string{
		"data_inicial": dataInicial.UTC().String(),
		"data_final":   dataFinal.UTC().String(),
	}

	var resp handler.TvceResponse
	err := s.Client.ListResource(RotaAudio+"/relatorio", resp, params)
	return resp, err
}

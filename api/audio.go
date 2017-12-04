package api

import (
	"strconv"

	"github.com/totalvoice/go-client/api/model"
)

// AudioService service
type AudioService struct {
	client  HTTPClient
	handler Response

	Relatorio *AudioRelatorioService
}

// NewAudioService - Serviço para o envio de Audio
func NewAudioService(httpClient HTTPClient, handler Response) *AudioService {

	service := &AudioService{
		client:  httpClient,
		handler: handler,
		Relatorio: &AudioRelatorioService{
			client:  httpClient,
			handler: handler,
		},
	}

	return service
}

// Enviar - Envia uma mensagem de audio
func (s AudioService) Enviar(numero string, urlAudio string, respostaUsuario bool, bina string) (*model.TotalVoiceResponse, error) {

	audio := new(model.Audio)
	audio.NumeroDestino = numero
	audio.URLAudio = urlAudio
	audio.RespostaUsuario = respostaUsuario
	audio.Bina = bina

	response := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(audio, RotaAudio)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

// BuscaAudio - Busca uma mensagem de audio pelo seu ID
func (s AudioService) BuscaAudio(id int) (*model.AudioResponse, error) {

	sID := strconv.Itoa(id)
	audio := new(model.Audio)
	response := new(model.AudioResponse)

	http, err := s.client.GetResource(audio, RotaAudio, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.AudioResponse), err
}

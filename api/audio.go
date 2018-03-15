package api

import (
	"strconv"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// AudioService service
type AudioService struct {
	client   HTTPClient
	response Response

	Relatorio *AudioRelatorioService
}

// NewAudioService - Serviço para o envio de Audio
func NewAudioService(httpClient HTTPClient, response Response) *AudioService {

	service := &AudioService{
		client:   httpClient,
		response: response,
		Relatorio: &AudioRelatorioService{
			client:   httpClient,
			response: response,
		},
	}

	return service
}

// Enviar - Envia uma mensagem de audio
func (s AudioService) Enviar(numero string, urlAudio string, respostaUsuario bool, bina string, gravaAudio bool) (*model.TotalVoiceResponse, error) {

	audio := new(model.Audio)
	audio.NumeroDestino = numero
	audio.URLAudio = urlAudio
	audio.RespostaUsuario = respostaUsuario
	audio.Bina = bina
	audio.GravarAudio = gravaAudio

	response := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(audio, RotaAudio)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

// Buscar - Busca uma mensagem de audio pelo seu ID
func (s AudioService) Buscar(id int) (*model.AudioResponse, error) {

	sID := strconv.Itoa(id)
	audio := new(model.Audio)
	response := new(model.AudioResponse)

	http, err := s.client.GetResource(audio, RotaAudio, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(response, http)
	return res.(*model.AudioResponse), err
}

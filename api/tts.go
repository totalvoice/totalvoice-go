package api

import (
	"strconv"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// TTSService service
type TTSService struct {
	client  HTTPClient
	handler Response

	Relatorio *TTSRelatorioService
}

// NewTTSService - Serviço para o envio de TTS
func NewTTSService(httpClient HTTPClient, handler Response) *TTSService {

	service := &TTSService{
		client:  httpClient,
		handler: handler,
		Relatorio: &TTSRelatorioService{
			client:  httpClient,
			handler: handler,
		},
	}

	return service
}

// Enviar - Envia uma mensagem de TTS
func (s TTSService) Enviar(numero string, mensagem string, opcoes map[string]interface{}) (*model.TotalVoiceResponse, error) {

	tts := new(model.TTS)
	tts.NumeroDestino = numero
	tts.Mensagem = mensagem

	for index, value := range opcoes {
		switch index {
		case "velocidade":
			tts.Velocidade = value.(int)
		case "resposta_usuario":
			tts.RespostaUsuario = value.(bool)
		case "tipo_voz":
			tts.TipoVoz = value.(string)
		case "bina":
			tts.Bina = value.(string)
		}
	}

	response := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(tts, RotaTTS)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

// Buscar - Busca uma mensagem de TTS pelo seu ID
func (s TTSService) Buscar(id int) (*model.TTSResponse, error) {

	sID := strconv.Itoa(id)
	TTS := new(model.TTS)
	response := new(model.TTSResponse)

	http, err := s.client.GetResource(TTS, RotaTTS, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TTSResponse), err
}

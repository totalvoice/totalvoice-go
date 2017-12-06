package api

import (
	"strconv"

	"github.com/totalvoice/go-client/api/model"
)

// SMSService service
type SMSService struct {
	client  HTTPClient
	handler Response

	Relatorio *SMSRelatorioService
}

// NewSMSService - Serviço para o envio de SMS
func NewSMSService(httpClient HTTPClient, handler Response) *SMSService {

	service := &SMSService{
		client:  httpClient,
		handler: handler,
		Relatorio: &SMSRelatorioService{
			client:  httpClient,
			handler: handler,
		},
	}

	return service
}

// Enviar - Envia uma mensagem de SMS
func (s SMSService) Enviar(numero string, mensagem string, respostaUsuario bool, multiSMS bool) (*model.TotalVoiceResponse, error) {

	sms := new(model.SMS)
	sms.NumeroDestino = numero
	sms.Mensagem = mensagem
	sms.RespostaUsuario = respostaUsuario
	sms.MultiSMS = multiSMS

	response := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(sms, RotaSMS)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

// BuscaSMS - Busca uma mensagem de SMS pelo seu ID
func (s SMSService) BuscaSMS(id int) (*model.SMSResponse, error) {

	sID := strconv.Itoa(id)
	SMS := new(model.SMS)
	response := new(model.SMSResponse)

	http, err := s.client.GetResource(SMS, RotaSMS, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.SMSResponse), err
}

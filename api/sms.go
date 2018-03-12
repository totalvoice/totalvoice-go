package api

import (
	"strconv"

	"github.com/totalvoice/totalvoice-go/api/model"
	"time"
)

// SMSService service
type SMSService struct {
	client   HTTPClient
	response Response

	Relatorio *SMSRelatorioService
}

// NewSMSService - Serviço para o envio de SMS
func NewSMSService(httpClient HTTPClient, response Response) *SMSService {

	service := &SMSService{
		client:   httpClient,
		response: response,
		Relatorio: &SMSRelatorioService{
			client:   httpClient,
			response: response,
		},
	}

	return service
}

// Enviar - Envia uma mensagem de SMS
func (s SMSService) Enviar(numero string, mensagem string, respostaUsuario bool, multiSMS bool, dataCriacao interface{}) (*model.TotalVoiceResponse, error) {

	sms := new(model.SMS)
	sms.NumeroDestino = numero
	sms.Mensagem = mensagem
	sms.RespostaUsuario = respostaUsuario
	sms.MultiSMS = multiSMS

	if dataCriacao != nil {

		data := dataCriacao.(time.Time)
		sms.DataCriacao = data.UTC().Format(DateFormat)
	}

	resp := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(sms, RotaSMS)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}

// Buscar - Busca uma mensagem de SMS pelo seu ID
func (s SMSService) Buscar(id int) (*model.SMSResponse, error) {

	sID := strconv.Itoa(id)
	SMS := new(model.SMS)
	resp := new(model.SMSResponse)

	http, err := s.client.GetResource(SMS, RotaSMS, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.SMSResponse), err
}

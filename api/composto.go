package api

import (
	"strconv"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// CompostoService service
type CompostoService struct {
	client  HTTPClient
	handler Response

	Relatorio *CompostoRelatorioService
}

// NewCompostoService - Serviço para mensagens Compostas
func NewCompostoService(httpClient HTTPClient, handler Response) *CompostoService {

	service := &CompostoService{
		client:  httpClient,
		handler: handler,
		Relatorio: &CompostoRelatorioService{
			client:  httpClient,
			handler: handler,
		},
	}

	return service
}

// Enviar - Envia uma mensagem de audio
func (s CompostoService) Enviar(composto model.Composto) (*model.TotalVoiceResponse, error) {

	resp := new(model.TotalVoiceResponse)
	http, err := s.client.CreateResource(composto, RotaComposto)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}

// Buscar - Busca uma mensagem composta pelo seu ID
func (s CompostoService) Buscar(id int) (*model.CompostoResponse, error) {

	sID := strconv.Itoa(id)
	composto := new(model.Composto)
	resp := new(model.CompostoResponse)

	http, err := s.client.GetResource(composto, RotaComposto, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(resp, http)
	return res.(*model.CompostoResponse), err
}

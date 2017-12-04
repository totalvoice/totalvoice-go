package api

import (
	"strconv"

	"github.com/totalvoice/go-client/api/model"
)

// CompostoService service
type CompostoService struct {
	client  HTTPClient
	handler Response

	Relatorio *CompostoRelatorioService
}

// NewCompostoService -
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
func (s CompostoService) Enviar(composto model.Composto) (*model.CompostoResponse, error) {

	resp := new(model.CompostoResponse)
	http, err := s.client.CreateResource(composto, RotaComposto)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(resp, http)
	return res.(*model.CompostoResponse), err
}

// BuscaComposto - Busca uma mensagem composta pelo seu ID
func (s CompostoService) BuscaComposto(id int) (*model.CompostoResponse, error) {

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

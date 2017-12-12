package api

import (
	"strconv"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// RamalService service
type RamalService struct {
	client  HTTPClient
	handler Response

	Relatorio *RamalRelatorioService
}

// NewRamalService - Serviço para o gerenciamento de Ramais
func NewRamalService(httpClient HTTPClient, handler Response) *RamalService {

	service := &RamalService{
		client:  httpClient,
		handler: handler,
		Relatorio: &RamalRelatorioService{
			client:  httpClient,
			handler: handler,
		},
	}

	return service
}

// Criar - Criar um ramal
func (s RamalService) Criar(ramal model.Ramal) (*model.RamalResponse, error) {

	response := new(model.RamalResponse)
	http, err := s.client.CreateResource(ramal, RotaRamal)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.RamalResponse), err
}

// Atualizar - Atualiza um ramal
func (s RamalService) Atualizar(ramal model.Ramal) (*model.RamalResponse, error) {

	sID := strconv.Itoa(ramal.ID)
	response := new(model.RamalResponse)
	http, err := s.client.UpdateResource(ramal, RotaRamal, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.RamalResponse), err
}

// Buscar - Busca uma mensagem de Ramal pelo seu ID
func (s RamalService) Buscar(id int) (*model.RamalResponse, error) {

	sID := strconv.Itoa(id)
	ramal := new(model.Ramal)
	response := new(model.RamalResponse)

	http, err := s.client.GetResource(ramal, RotaRamal, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.RamalResponse), err
}

// Excluir - Remove um Ramal
func (s RamalService) Excluir(id int) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	response := new(model.TotalVoiceResponse)

	http, err := s.client.DeleteResource(RotaRamal, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

// Webphone - Requisita a URL do Webphone de um ramal
func (s RamalService) Webphone(params map[string]interface{}) (*model.WebphoneResponse, error) {

	webphone := new(model.Webphone)
	response := new(model.WebphoneResponse)

	http, err := s.client.ListResource(webphone, RotaWebphone, params)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.WebphoneResponse), err
}

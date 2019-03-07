package api

import (
	"strconv"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// ValidaNumeroService service
type ValidaNumeroService struct {
	client   HTTPClient
	response Response

	Relatorio *ValidaNumeroRelatorioService
}

// NewValidaNumeroService - Serviço para ValidaNunero
func NewValidaNumeroService(httpClient HTTPClient, response Response) *ValidaNumeroService {

	service := &ValidaNumeroService {
		client:   httpClient,
		response: response,
		Relatorio: &ValidaNumeroRelatorioService {
			client:   httpClient,
			response: response,
		},
	}
	
	return service
}

// Criar - Envia uma mensagem de ValidaNumero
func (s ValidaNumeroService) Criar(numeroDestino string) (*model.ValidaNumeroResponse, error) {

	validaNumero := new (model.ValidaNumero)
	validaNumero.NumeroDestino = numeroDestino
	resp := new(model.ValidaNumeroResponse)
	http, err := s.client.CreateResource(validaNumero, RotaValidaNumero)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.ValidaNumeroResponse), err
}

// Buscar - Busca um ValidaNumero pelo seu ID
func (s ValidaNumeroService) Buscar(id int) (*model.ValidaNumeroResponse, error) {

	sID := strconv.Itoa(id)
	resp := new(model.ValidaNumeroResponse)

	http, err := s.client.GetResource(nil, RotaValidaNumero, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.ValidaNumeroResponse), err
}
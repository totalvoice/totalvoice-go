package api

import (
	"strconv"

	"github.com/totalvoice/go-client/api/model"
)

// URAService service
type URAService struct {
	client  HTTPClient
	handler Response
}

// NewURAService - Serviço para o gerenciamento de URAs
func NewURAService(httpClient HTTPClient, handler Response) *URAService {

	service := &URAService{
		client:  httpClient,
		handler: handler,
	}

	return service
}

// Criar - Criar uma URA
func (s URAService) Criar(ura model.URA) (*model.TotalVoiceResponse, error) {

	response := new(model.TotalVoiceResponse)
	http, err := s.client.CreateResource(ura, RotaURA)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

// Atualizar - Atualiza a URA
func (s URAService) Atualizar(ura model.URA) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(ura.ID)
	response := new(model.TotalVoiceResponse)
	http, err := s.client.UpdateResource(ura, RotaURA, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

// Excluir - Remove uma URA
func (s URAService) Excluir(id int) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	response := new(model.TotalVoiceResponse)

	http, err := s.client.DeleteResource(RotaURA, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

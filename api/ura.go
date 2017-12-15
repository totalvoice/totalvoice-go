package api

import (
	"strconv"

	"github.com/totalvoice/totalvoice-go/api/model"
)

// URAService service
type URAService struct {
	client   HTTPClient
	response Response
}

// NewURAService - Serviço para o gerenciamento de URAs
func NewURAService(httpClient HTTPClient, response Response) *URAService {

	service := &URAService{
		client:   httpClient,
		response: response,
	}

	return service
}

// Criar - Criar uma URA
func (s URAService) Criar(ura model.URA) (*model.TotalVoiceResponse, error) {

	resp := new(model.TotalVoiceResponse)
	http, err := s.client.CreateResource(ura, RotaURA)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}

// Atualizar - Atualiza a URA
func (s URAService) Atualizar(ura model.URA) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(ura.ID)
	resp := new(model.TotalVoiceResponse)
	http, err := s.client.UpdateResource(ura, RotaURA, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}

// Excluir - Remove uma URA
func (s URAService) Excluir(id int) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	resp := new(model.TotalVoiceResponse)

	http, err := s.client.DeleteResource(RotaURA, sID)
	if err != nil {
		return nil, err
	}
	res := s.response.HandleResponse(resp, http)
	return res.(*model.TotalVoiceResponse), err
}

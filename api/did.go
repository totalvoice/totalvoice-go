package api

import (
	"strconv"

	"github.com/totalvoice/go-client/api/model"
)

// DIDService service
type DIDService struct {
	client  HTTPClient
	handler Response
}

// NewDIDService - Serviço para o gerenciamento de Ramais
func NewDIDService(httpClient HTTPClient, handler Response) *DIDService {

	service := &DIDService{
		client:  httpClient,
		handler: handler,
	}

	return service
}

// Atualizar - Atualiza os dados de um DID
func (s DIDService) Atualizar(did model.DID) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(did.ID)
	response := new(model.TotalVoiceResponse)
	http, err := s.client.UpdateResource(did, RotaDID, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

// Listar - Busca uma mensagem de DID pelo seu ID
func (s DIDService) Listar() (*model.DIDResponse, error) {

	did := new(model.DID)
	response := new(model.DIDResponse)

	http, err := s.client.ListResource(did, RotaDID, nil)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.DIDResponse), err
}

// Excluir - Remove um DID
func (s DIDService) Excluir(id int) (*model.TotalVoiceResponse, error) {

	sID := strconv.Itoa(id)
	response := new(model.TotalVoiceResponse)

	http, err := s.client.DeleteResource(RotaDID, sID)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}

// Estoque - Consulta a lista de DIDs disponiveis
func (s DIDService) Estoque() (*model.DIDResponse, error) {

	did := new(model.DID)
	response := new(model.DIDResponse)

	http, err := s.client.ListResource(did, RotaDID+"/estoque", nil)
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.DIDResponse), err
}

// Adquirir - Adquire um DID para a Conta
func (s DIDService) Adquirir(id int) (*model.TotalVoiceResponse, error) {

	did := new(model.DID)
	did.DidID = strconv.Itoa(id)
	response := new(model.TotalVoiceResponse)

	http, err := s.client.CreateResource(did, RotaDID+"/estoque")
	if err != nil {
		return nil, err
	}
	res := s.handler.HandleResponse(response, http)
	return res.(*model.TotalVoiceResponse), err
}
